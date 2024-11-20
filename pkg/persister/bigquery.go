package persister

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/nais/dataproduct-topics/pkg/collector"
)

const BigQueryProjectID = "nais-prod-b6f2"

func Persist(ctx context.Context, topics []collector.Topic) error {
	client, err := bigquery.NewClient(ctx, BigQueryProjectID)
	if err != nil {
		return err
	}

	tableHandle := client.Dataset("dataproduct_topics").Table("dataproduct_topics")
	inserter := tableHandle.Inserter()
	return inserter.Put(ctx, topics)
}
