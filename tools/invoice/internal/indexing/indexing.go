package indexing

import (
	"context"
	"fmt"
	"log"
	"time"

	"invoice/internal/duration"

	"github.com/meilisearch/meilisearch-go"
)

const Timeout = time.Second * 120
const Index = "invoice"

func RunToInvoice(client *meilisearch.Client, documents []map[string]interface{}) error {
	resp, err := client.Index(Index).AddDocuments(documents, "registratedNumber")
	if err != nil {
		return fmt.Errorf("fail to start AddDocuments: %w", err)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), Timeout)
	defer cancelFunc()

	task, err := client.WaitForTask(resp.TaskUID, meilisearch.WaitParams{
		Context:  ctx,
		Interval: time.Second * 3,
	})
	if err != nil {
		return fmt.Errorf("fail to waiting AddDocuments: %w", err)
	}

	// MEMO: 所要時間を出力しているがIndexingTimeoutが適切に決まれば無くしても良い
	duration, err := duration.ParseDurationISO8061(task.Duration)
	if err != nil {
		return err
	}
	log.Printf("success AddDocuments: duration %s (%v)", duration.String(), task.Duration)
	return nil
}
