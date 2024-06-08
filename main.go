package main

import (
	"api-with-go-postgres/database"
	"api-with-go-postgres/products"
	"log"
	"net/http"
)

func main() {
	db := database.Connect()

	// Route to create a new product
	http.HandleFunc("/products/create", products.CreateProduct(db))

	// Starting the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor HTTP:", err)
	}
}
