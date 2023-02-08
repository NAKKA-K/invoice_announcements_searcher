package main

import (
	"encoding/json"
	"github.com/meilisearch/meilisearch-go"
	"log"
	"sync"
)

const DataDir = "./data"
const Index = "invoice"

func LoadJSON(filename string, client *meilisearch.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	announcements, err := LoadFromJson(filename)
	if err != nil {
		return
	}
	// documents := structs.Map(announcements)
	// struct to map
	var documents []map[string]interface{}
	b, _ := json.Marshal(announcements)
	json.Unmarshal(b, &documents)

	err = RunIndexingInvoice(client, documents)
	if err != nil {
		log.Fatalf("loaded: %s, err: %v", filename, err)
	} else {
		log.Printf("loaded: %s, status: SUCCESS", filename)
	}
}

func main() {
	files, err := GetFileNames(DataDir)
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
		go LoadJSON(DataDir+"/"+file, client, &wg)
	}
	wg.Wait()

	log.Println("finish loading json")
}
