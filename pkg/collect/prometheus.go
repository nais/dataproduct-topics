package collect

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

func GetKafkaTopicsFromPrometheus(ctx context.Context, url string) ([]string, error) {
	cfg := api.Config{
		Address:      url,
		RoundTripper: nil,
	}
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	apiClient := v1.NewAPI(client)
	value, warnings, err := apiClient.Query(ctx, "kafka_server_brokertopicmetrics_messagesin_total", time.Now())
	if err != nil {
		return nil, err
	}

	for _, warning := range warnings {
		log.Warnf("%v", warning)
	}

	vec := value.(model.Vector)
	topics := make([]string, len(vec))
	for i, row := range vec {
		topics[i] = fmt.Sprintf("%s", row.Metric["topic"])
	}

	return topics, nil
}
