package product

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)
type Repository struct {
	es *elasticsearch.Client
}


// NewRepository gives the employee a phone line to the warehouse
func NewRepository() *Repository {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{
		es: es,
	}
}


func (r *Repository) Search(
	ctx context.Context,
	query map[string]interface{},
) ([]Product, error) {

	// Convert instructions to JSON (notebook language)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	// Ask Elasticsearch to READ the notebook
	res, err := r.es.Search(
		r.es.Search.WithContext(ctx),
		r.es.Search.WithIndex("products"), // products notebook
		r.es.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// If notebook replies with error
	if res.IsError() {
		return nil, errors.New("elasticsearch search error")
	}

	// Structure matching Elasticsearch response
	var esResp struct {
		Hits struct {
			Hits []struct {
				Source Product `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	// Read notebook pages
	if err := json.NewDecoder(res.Body).Decode(&esResp); err != nil {
		return nil, err
	}

	// Extract clean product list
	products := make([]Product, 0, len(esResp.Hits.Hits))
	for _, hit := range esResp.Hits.Hits {
		products = append(products, hit.Source)
	}

	return products, nil
}
