package main

import (
	"api-with-go-postgres/auth"
	"api-with-go-postgres/database"
	"api-with-go-postgres/products"
)
func main(){
	database.Connect()
	auth.Login()
	auth.ForgotPassword()
	auth.ResetPassword()
	products.CreateProduct()
	products.DeleteProductById()
	products.GetAllProducts()
	products.UpdateProductById()
}