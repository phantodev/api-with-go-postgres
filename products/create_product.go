package products

import (
	"api-with-go-postgres/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a variable to store the data from the request body
		var product models.Products
		// Decode the data from the request body into the product variable
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, "Error to decode the request", http.StatusBadRequest)
			return
		}

		// Execute SQL query to insert the product into the database
		_, err = db.Exec("INSERT INTO products (name, description) VALUES ($1, $2)", product.Name, product.Description)
		if err != nil {
			http.Error(w, "Error to insert Product in database", http.StatusInternalServerError)
			return
		}

		// Response with status code 201 (Created) and a message indicating that the product was inserted successfully
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Product created in database"))
	}
}
