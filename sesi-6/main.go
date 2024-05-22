package main

import (
	"golang-for-women/sesi-6/config"
	"golang-for-women/sesi-6/controllers"
	"log"

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

	// Products
	// controllers.CreateProduct("Meja makan")
	// controllers.UpdateProduct(1, "Side table")
	// controllers.GetProductById(2)

	// Variants
	// variantData := models.Variant{
	// 	Name:      "Kayu jati",
	// 	Quantity:  4,
	// 	ProductID: 11,
	// }
	// controllers.CreateVariant(variantData)
	// controllers.UpdateVariantById(1, variantData)
	// controllers.DeleteVariantById(2)

	controllers.GetProductWithVariant()
}
