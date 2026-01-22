package product

import "context"

// Service decides HOW search should happen
// Amazon analogy: decision maker for product search
type Service struct {
	repo *Repository
}

// NewService injects repository
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// SearchProducts builds Elasticsearch query and asks repository to execute it
func (s *Service) SearchProducts(
	ctx context.Context,
	text string,
	brand string,
	category string,
	minPrice float64,
	maxPrice float64,
	inStock *bool,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]Product, error) {

	must := make([]interface{}, 0)

	if text != "" {
		must = append(must, map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  text,
				"fields": []string{"name", "description"},
			},
		})
	}

	filters := make([]interface{}, 0)

	if brand != "" {
		filters = append(filters, map[string]interface{}{
			"term": map[string]interface{}{
				"brand": brand,
			},
		})
	}

	if category != "" {
		filters = append(filters, map[string]interface{}{
			"term": map[string]interface{}{
				"category": category,
			},
		})
	}

	if minPrice > 0 || maxPrice > 0 {
		priceRange := map[string]interface{}{}
		if minPrice > 0 {
			priceRange["gte"] = minPrice
		}
		if maxPrice > 0 {
			priceRange["lte"] = maxPrice
		}

		filters = append(filters, map[string]interface{}{
			"range": map[string]interface{}{
				"price": priceRange,
			},
		})
	}

	if inStock != nil {
		filters = append(filters, map[string]interface{}{
			"term": map[string]interface{}{
				"in_stock": *inStock,
			},
		})
	}

	from := (page - 1) * limit
	if from < 0 {
		from = 0
	}

	sort := make([]interface{}, 0)
	if sortBy != "" {
		sort = append(sort, map[string]interface{}{
			sortBy: map[string]interface{}{
				"order": sortOrder,
			},
		})
	}

	query := map[string]interface{}{
		"from": from,
		"size": limit,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must":   must,
				"filter": filters,
			},
		},
	}

	if len(sort) > 0 {
		query["sort"] = sort
	}

	return s.repo.Search(ctx, query)
}
