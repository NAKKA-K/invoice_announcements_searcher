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
	fileName string
}

type ChResp string

func Run(req *Request) {
	defer req.wg.Done()

	announcements, err := announcement.LoadFromJson(req.fileName)
	if err != nil {
		req.ch <- ChResp(err.Error())
		return
	}
	documents, err := announcement.ToMaps(announcements)
	if err != nil {
		req.ch <- ChResp(err.Error())
		return
	}

	dur, err := indexing.ToInvoice(req.client, documents)
	if err != nil {
		req.ch <- ChResp(fmt.Sprintf("loaded: %s, err: %v", req.fileName, err))
	} else {
		req.ch <- ChResp(fmt.Sprintf("loaded: %s, status: SUCCESS, duration %s", req.fileName, dur.String()))
	}
}

func NewRequest(wg *sync.WaitGroup, ch chan<- ChResp, client *meilisearch.Client, fileName string) *Request {
	return &Request{
		wg:       wg,
		ch:       ch,
		client:   client,
		fileName: fileName,
	}
}
