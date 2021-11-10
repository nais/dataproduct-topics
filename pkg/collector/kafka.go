package collector

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	log "github.com/sirupsen/logrus"
)

type Topic struct {
	CollectionTime time.Time `bigquery:"collection_time"`
	Topic          string    `bigquery:"topic"`
	Team           string    `bigquery:"team"`
	Cluster        string    `bigquery:"cluster"`
}

type Collector struct {
	dialer *kafka.Dialer
}

func (c *Collector) ConfigureOnpremDialer() error {
	mechanism := plain.Mechanism{
		Username: os.Getenv("KAFKA_USERNAME"),
		Password: os.Getenv("KAFKA_PASSWORD"),
	}

	certPool, err := x509.SystemCertPool()
	if err != nil {
		return fmt.Errorf("load system cert pool: %s", err)
	}

	c.dialer = &kafka.Dialer{
		DualStack:     false,
		SASLMechanism: mechanism,
		TLS: &tls.Config{
			RootCAs: certPool,
		},
	}

	return nil
}

func (c *Collector) ConfigureAivenDialer() error {
	ca := os.Getenv("KAFKA_CA")
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM([]byte(ca))

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

	cluster := os.Getenv("CLUSTER_NAME")

	topicList := make([]Topic, len(topicMap))
	i := 0
	for key := range topicMap {
		topicList[i] = createTopicFromName(key, cluster)
		i++
	}

	return topicList, nil
}

func createTopicFromName(topicName, cluster string) Topic {
	parts := strings.SplitN(topicName, ".", 2)

	if len(parts) == 2 {
		return Topic{
			CollectionTime: time.Now(),
			Topic:          parts[1],
			Team:           parts[0],
			Cluster:        cluster,
		}
	}

	return Topic{
		CollectionTime: time.Now(),
		Topic:          topicName,
		Team:           "",
		Cluster:        cluster,
	}
}
