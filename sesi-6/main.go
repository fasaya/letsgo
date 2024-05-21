package main

import (
	"fmt"
	"golang-for-women/sesi-6/config"
	product_controller "golang-for-women/sesi-6/controllers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var PORT = ":9090"

func main() {

	// Initialize the database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Set the database connection in the config package
	config.DB = db

	// fmt.Println("Successfully connected to database", db_name)

	http.HandleFunc("/", product_controller.Index)
	http.HandleFunc("/products", product_controller.GetProductWithVariant)
	http.HandleFunc("/product/update/", product_controller.UpdateProduct)
	http.HandleFunc("/product/", product_controller.GetProductById)
	http.HandleFunc("/product", product_controller.CreateProduct)

	fmt.Println("Application is listening on port", PORT, "and successfully connected to database.")
	http.ListenAndServe(PORT, nil)

	// createProduct()
	// updateProduct()
	// getProductById()
	// createVariant()
	// updateVariantById()
	// deleteVariantById()
	// getProductWithVariant()
}
