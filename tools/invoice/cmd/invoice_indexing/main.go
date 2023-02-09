package main

import (
	"log"
	"sync"

	"invoice/internal/file"
	"invoice/internal/importer"

	"github.com/meilisearch/meilisearch-go"
)

const DataDir = "./data"

func main() {
	files, err := file.GetFileNames(DataDir)
	if err != nil {
		log.Fatal(err)
	}

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://localhost:7700",
	})

	log.Println("start to load json")

	var wg sync.WaitGroup
	ch := make(chan importer.ChResp, len(files))
	for _, file := range files {
		wg.Add(1)

		req := importer.NewRequest(&wg, ch, client, DataDir+"/"+file)
		go importer.Run(req)
	}

	for i := 0; i < len(files); i++ {
		msg := <-ch
		log.Println(string(msg))
	}

	close(ch)
	wg.Wait()

	log.Println("finish loading json")
}
