package importer

import (
	"fmt"
	"sync"

	"invoice/internal/announcement"
	"invoice/internal/indexing"

	"github.com/meilisearch/meilisearch-go"
)

type Request struct {
	wg       *sync.WaitGroup
	ch       chan<- ChResp
	client   *meilisearch.Client
	fineName string
}

type ChResp string

func LoadJSON(req *Request) {
	defer req.wg.Done()

	announcements, err := announcement.LoadFromJson(req.fineName)
	if err != nil {
		req.ch <- ChResp(err.Error())
		return
	}
	documents, err := announcement.ToSliceMap(announcements)
	if err != nil {
		req.ch <- ChResp(err.Error())
		return
	}

	dur, err := indexing.RunToInvoice(req.client, documents)
	if err != nil {
		req.ch <- ChResp(fmt.Sprintf("loaded: %s, err: %v", req.fineName, err))
	} else {
		req.ch <- ChResp(fmt.Sprintf("loaded: %s, status: SUCCESS, duration %s", req.fineName, dur.String()))
	}
}

func NewRequest(wg *sync.WaitGroup, ch chan<- ChResp, client *meilisearch.Client, fileName string) *Request {
	return &Request{
		wg:       wg,
		ch:       ch,
		client:   client,
		fineName: fileName,
	}
}
