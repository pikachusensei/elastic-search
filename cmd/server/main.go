package main

import (
	"log"
	"net/http"

	"product-search/internal/product"
	"product-search/internal/server"
)

func main() {
	repo := product.NewRepository()

	service := product.NewService(repo)

	handler := product.NewHandler(service)

	router := server.NewRouter(handler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
