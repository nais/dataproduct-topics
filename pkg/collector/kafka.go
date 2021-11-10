package collector

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	log "github.com/sirupsen/logrus"
)

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

func (c *Collector) GetTopics(ctx context.Context, brokers []string) ([]string, error) {
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

	topicList := make([]string, len(topicMap))
	i := 0
	for key := range topicMap {
		topicList[i] = key
		i++
	}

	return topicList, nil
}
