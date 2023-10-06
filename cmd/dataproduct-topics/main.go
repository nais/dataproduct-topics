package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nais/dataproduct-topics/pkg/collector"
	"github.com/nais/dataproduct-topics/pkg/persister"
	log "github.com/sirupsen/logrus"
)

func main() {
	programContext, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	topics, err := collect(programContext)
	if err != nil {
		log.Errorf("getting topics: %s", err)
		os.Exit(1)
	}

	err = persister.Persist(programContext, topics)
	if err != nil {
		log.Errorf("persisting topics: %s", err)
		os.Exit(1)
	}

	log.Info("Done!")
}

func collect(ctx context.Context) ([]collector.Topic, error) {
	collect := &collector.Collector{}
	err := collect.ConfigureAivenDialer()
	if err != nil {
		return nil, fmt.Errorf("configure aiven collector: %s", err)
	}

	brokersEnv := os.Getenv("KAFKA_BROKERS")
	brokers := strings.Split(brokersEnv, ",")
	return collect.GetTopics(ctx, brokers)
}
