package main

import (
	"context"
	"os"
	"strings"

	"github.com/nais/dataproduct-topics/pkg/collector"
	log "github.com/sirupsen/logrus"
)

func main() {
	programContext := context.Background()

	topics, err := collect(programContext)
	if err != nil {
		log.Errorf("getting topics: %s", err)
	}

	for _, topic := range topics {
		log.Info(topic)
	}
}

func collect(ctx context.Context) ([]string, error) {
	collect := &collector.Collector{}
	if os.Getenv("ONPREM") == "true" {
		err := collect.ConfigureOnpremDialer()
		if err != nil {
			log.Errorf("Setting up onprem collector: %s", err)
			os.Exit(1)
		}
	} else {
		err := collect.ConfigureAivenDialer()
		if err != nil {
			log.Errorf("Setting up aiven collector: %s", err)
			os.Exit(1)
		}
	}

	brokersEnv := os.Getenv("KAFKA_BROKERS")
	brokers := strings.Split(brokersEnv, ",")
	return collect.GetTopics(ctx, brokers)
}
