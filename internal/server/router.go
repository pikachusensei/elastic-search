package server

import (
	"net/http"

	"product-search/internal/product"
)

func NewRouter(handler *product.Handler) http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/products/search", handler.SearchProducts)

	return mux
}
