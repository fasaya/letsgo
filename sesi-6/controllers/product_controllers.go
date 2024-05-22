package controllers

import (
	"fmt"
	"golang-for-women/sesi-6/config"
	"golang-for-women/sesi-6/models"
	"time"
)

func GetProductWithVariant() {

	var products = []models.Product{}

	sqlStatement := `SELECT * from products`

	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product = models.Product{}

		err = rows.Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	fmt.Printf("DAll data: %+v\n", products)
}

func CreateProduct(productName string) {
	var product = models.Product{}

	// You can further process the request body data as needed
	currentTime := time.Now()

	sqlStatement := `INSERT INTO products (name, created_at, updated_at) VALUES (?, ?, ?)`

	result, err := config.DB.Exec(sqlStatement, productName, currentTime, currentTime)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Retrieve the inserted row using the lastInsertID
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	err = config.DB.QueryRow(sqlRetrieve, lastInsertID).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data created: %+v\n", product)

}

func UpdateProduct(productId int, productName string) {
	var product = models.Product{}

	// You can further process the request body data as needed
	currentTime := time.Now()

	sqlStatement := `UPDATE products SET name = ?, updated_at = ? WHERE id = ?;`

	result, err := config.DB.Exec(sqlStatement, productName, currentTime, productId)
	if err != nil {
		panic(err)
	}

	// Get affected rows
	_, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}
	// fmt.Println("Rows affected:", count, "row(s)")

	// Retrieve the inserted row using the id
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	err = config.DB.QueryRow(sqlRetrieve, productId).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data updated: %+v\n", product)
}

func GetProductById(productId int) {

	var product = models.Product{}

	// Retrieve the inserted row using the id
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	err := config.DB.QueryRow(sqlRetrieve, productId).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data by id: %+v\n", product)

}
