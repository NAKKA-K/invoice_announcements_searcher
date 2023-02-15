package importer

import (
	"context"
	"log"
	"sync"

	"invoice/internal/announcement"
	"invoice/internal/indexing"

	"github.com/meilisearch/meilisearch-go"
)

type Request struct {
	wg       *sync.WaitGroup
	ctx      context.Context
	client   *meilisearch.Client
	fileName string
}

func Run(req *Request) {
	defer req.wg.Done()

	announcements, err := announcement.LoadFromJson(req.fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	documents, err := announcement.ToJSONMaps(announcements)
	if err != nil {
		log.Fatal(err)
		return
	}

	dur, err := indexing.ToInvoice(req.ctx, req.client, documents)
	if err != nil {
		log.Fatalf("loaded: %s, err: %v", req.fileName, err)
	} else {
		log.Printf("loaded: %s, status: SUCCESS, duration %s", req.fileName, dur.String())
	}
}

func NewRequest(wg *sync.WaitGroup, ctx context.Context, client *meilisearch.Client, fileName string) *Request {
	return &Request{
		wg:       wg,
		ctx:      ctx,
		client:   client,
		fileName: fileName,
	}
}
