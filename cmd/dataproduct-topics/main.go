package main

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/nais/dataproduct-topics/pkg/collector"
	log "github.com/sirupsen/logrus"
)

var (
	onprem  = false
)

func init() {
	flag.BoolVar(&onprem, "onprem", true, "use onprem settings")
	flag.Parse()
}

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
	if onprem {
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

	return collect.GetTopics(ctx, []string{"b27apvl00045.preprod.local:8443"})
}
