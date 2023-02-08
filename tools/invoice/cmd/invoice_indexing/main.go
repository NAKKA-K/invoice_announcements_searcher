package main

import (
	"invoice/internal/file"
	"invoice/internal/importer"
	"log"
	"sync"

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
	for _, file := range files {
		wg.Add(1)
		go importer.LoadJSON(DataDir+"/"+file, client, &wg)
	}
	wg.Wait()

	log.Println("finish loading json")
}
