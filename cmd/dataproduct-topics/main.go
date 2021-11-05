package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/nais/dataproduct-topics/pkg/collect"
)

func main() {
	programContext := context.Background()

	prometheusTopics, err := collect.GetKafkaTopicsFromPrometheus(programContext, "https://prometheus.dev-fss.nais.io")
	if err != nil {
		log.Errorf("collect kafka topics: %s", err)
		os.Exit(1)
	}

	adminRestTopics, err := collect.GetKafkaTopicsFromKafkaAdminRest(programContext, "https://kafka-adminrest.dev-fss.nais.io")
	if err != nil {
		log.Errorf("collect kafka topics: %s", err)
		os.Exit(1)
	}

	//fmt.Println(prometheusTopics)
	//fmt.Println(adminRestTopics)

	//outer:
	//for _, restTopic := range adminRestTopics {
	//	for _, promTopic := range prometheusTopics {
	//		if restTopic == promTopic {
	//			continue outer
	//		}
	//	}
	//	fmt.Printf("%s is in rest, but not prom\n", restTopic)
	//}
	//
	//outer2:
	//for _, promTopic := range prometheusTopics {
	//	for _, restTopic := range adminRestTopics {
	//		if restTopic == promTopic {
	//			continue outer2
	//		}
	//	}
	//	fmt.Printf("%s is in prom, but not rest\n", promTopic)
	//}

	adminRestTopics = sanitize(adminRestTopics)
	prometheusTopics = sanitize(prometheusTopics)

	fmt.Printf("len(adminRestTopics)=%d\n", len(adminRestTopics))
	fmt.Printf("len(prometheusTopics)=%d\n", len(prometheusTopics))
}

func sanitize(topics []string) []string {
	topicsMap := make(map[string]struct{})
	for _, topic := range topics {
		trimmed := strings.Trim(topic, " ")
		if trimmed == "" {
			continue
		}

		topicsMap[trimmed] = struct{}{}
	}

	keys := make([]string, len(topicsMap))
	i := 0
	for topic := range topicsMap {
		keys[i] = topic
		i++
	}

	sort.Strings(keys)

	return keys
}
