package importer

import (
	"log"
	"sync"

	"invoice/internal/announcement"
	"invoice/internal/indexing"

	"github.com/meilisearch/meilisearch-go"
)

func LoadJSON(filename string, client *meilisearch.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	announcements, err := announcement.LoadFromJson(filename)
	if err != nil {
		return
	}
	documents := announcement.ToSliceMap(announcements)

	err = indexing.RunToInvoice(client, documents)
	if err != nil {
		log.Fatalf("loaded: %s, err: %v", filename, err)
	} else {
		log.Printf("loaded: %s, status: SUCCESS", filename)
	}
}
