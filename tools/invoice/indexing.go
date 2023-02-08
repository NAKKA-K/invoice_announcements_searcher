package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/meilisearch/meilisearch-go"
	"github.com/sosodev/duration"
)

const IndexingTimeout = time.Second * 60
const Index = "invoice"

func RunIndexingInvoice(client *meilisearch.Client, documents []map[string]interface{}) error {
	resp, err := client.Index(Index).AddDocuments(documents, "registratedNumber")
	if err != nil {
		return fmt.Errorf("fail to start AddDocuments: %w", err)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), IndexingTimeout)
	defer cancelFunc()

	task, err := client.WaitForTask(resp.TaskUID, meilisearch.WaitParams{
		Context:  ctx,
		Interval: time.Second * 3,
	})
	if err != nil {
		return fmt.Errorf("fail to waiting AddDocuments: %w", err)
	}

	// MEMO: 所要時間を出力しているがIndexingTimeoutが適切に決まれば無くしても良い
	duration, err := parseDurationISO8061(task.Duration)
	if err != nil {
		return err
	}
	log.Printf("success AddDocuments: duration %s", duration.String())
	return nil
}

func parseDurationISO8061(meiliDuration string) (*time.Duration, error) {
	d, err := duration.Parse(meiliDuration)
	if err != nil {
		return nil, err
	}

	duration := d.ToTimeDuration()
	return &duration, nil
}
