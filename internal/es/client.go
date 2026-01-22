package es

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewClient() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
