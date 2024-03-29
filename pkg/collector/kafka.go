package collector

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type Topic struct {
	CollectionTime civil.DateTime `bigquery:"collection_time"`
	Topic          string         `bigquery:"topic"`
	Team           string         `bigquery:"team"`
	Pool           string         `bigquery:"pool"`
}

type Collector struct {
	dialer *kafka.Dialer
}

var topicFilters []*regexp.Regexp

func init() {
	topicFilters = []*regexp.Regexp{
		// Kafka streams topics are linked to source/destination topics
		regexp.MustCompile("(?i).*?-.*store.*-(changelog|repartition)"),
		regexp.MustCompile("(?i).*?streams-[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}-.*"),
		regexp.MustCompile(".*ProcessingEventDtos-changelog"),
		// Built-in or semi-built-in topics aren't that interesting
		regexp.MustCompile("__.*"),
		regexp.MustCompile("_schemas"),
	}
}

func (c *Collector) ConfigureAivenDialer() error {
	pool := x509.NewCertPool()
	ca := os.Getenv("KAFKA_CA")
	if len(ca) > 0 {
		pool.AppendCertsFromPEM([]byte(ca))
	} else {
		caPath := os.Getenv("KAFKA_CA_PATH")
		caBytes, err := os.ReadFile(caPath)
		if err != nil {
			return fmt.Errorf("unable to read CA: %w", err)
		}
		pool.AppendCertsFromPEM(caBytes)
	}

	certFile := os.Getenv("KAFKA_CERTIFICATE_PATH")
	keyFile := os.Getenv("KAFKA_PRIVATE_KEY_PATH")

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	c.dialer = &kafka.Dialer{
		TLS: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cert},
		},
	}
	log.Info("configured Aiven dialer")

	return nil
}

func (c *Collector) connect(ctx context.Context, brokers []string) (*kafka.Conn, error) {
	for _, broker := range brokers {
		client, err := c.dialer.DialContext(ctx, "tcp", broker)
		if err == nil {
			return client, nil
		}
		log.Warnf("dialing broker %s: %s", broker, err)
	}

	return nil, fmt.Errorf("failed connecting to any brokers")
}

func (c *Collector) GetTopics(ctx context.Context, brokers []string) ([]Topic, error) {
	client, err := c.connect(ctx, brokers)
	if err != nil {
		return nil, err
	}

	partitions, err := client.ReadPartitions()
	if err != nil {
		return nil, fmt.Errorf("read partition: %s", err)
	}

	topicMap := make(map[string]struct{}, len(partitions))
	for _, partition := range partitions {
		topicMap[partition.Topic] = struct{}{}
	}

	pool := os.Getenv("POOL_NAME")

	topicList := make([]Topic, 0, len(topicMap))
	log.Infof("found %d topics in %s", len(topicMap), pool)
	for key := range topicMap {
		if ignoreTopic(key) {
			continue
		}
		topic, err := createTopicFromName(key, pool)
		if err != nil {
			log.Warn(err)
			continue
		}
		topicList = append(topicList, topic)
	}
	log.Infof("found %d interesting topics in %s", len(topicList), pool)

	return topicList, nil
}

func ignoreTopic(topicName string) bool {
	for _, filter := range topicFilters {
		if filter.MatchString(topicName) {
			return true
		}
	}
	return false
}

func createTopicFromName(topicName, pool string) (Topic, error) {
	parts := strings.SplitN(topicName, ".", 2)

	if len(parts) == 2 {
		return Topic{
			CollectionTime: civil.DateTimeOf(time.Now()),
			Topic:          parts[1],
			Team:           parts[0],
			Pool:           pool,
		}, nil
	}

	return Topic{}, fmt.Errorf("unable to parse topic name: %s", topicName)
}
