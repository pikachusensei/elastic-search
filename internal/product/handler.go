package product

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SearchProducts(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	text := q.Get("q")
	brand := q.Get("brand")
	category := q.Get("category")
	sortBy := q.Get("sort_by")
	sortOrder := q.Get("sort_order")

	minPrice, _ := strconv.ParseFloat(q.Get("min_price"), 64)
	maxPrice, _ := strconv.ParseFloat(q.Get("max_price"), 64)

	page, _ := strconv.Atoi(q.Get("page"))
	limit, _ := strconv.Atoi(q.Get("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	var inStock *bool
	if q.Get("in_stock") != "" {
		val, _ := strconv.ParseBool(q.Get("in_stock"))
		inStock = &val
	}

	products, err := h.service.SearchProducts(
		r.Context(),
		text,
		brand,
		category,
		minPrice,
		maxPrice,
		inStock,
		page,
		limit,
		sortBy,
		sortOrder,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
