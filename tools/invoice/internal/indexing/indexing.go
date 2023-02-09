package indexing

import (
	"context"
	"fmt"
	"time"

	"invoice/internal/duration"

	"github.com/meilisearch/meilisearch-go"
)

const Timeout = time.Second * 120
const Index = "invoice"

func ToInvoice(client *meilisearch.Client, documents []map[string]interface{}) (*time.Duration, error) {
	resp, err := client.Index(Index).AddDocuments(documents, "registratedNumber")
	if err != nil {
		return nil, fmt.Errorf("fail to start AddDocuments: %w", err)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), Timeout)
	defer cancelFunc()

	task, err := client.WaitForTask(resp.TaskUID, meilisearch.WaitParams{
		Context:  ctx,
		Interval: time.Second * 3,
	})
	if err != nil {
		return nil, fmt.Errorf("fail to waiting AddDocuments: %w", err)
	}

	dur, err := duration.ParseDurationISO8061(task.Duration)
	if err != nil {
		return nil, err
	}
	return dur, nil
}
