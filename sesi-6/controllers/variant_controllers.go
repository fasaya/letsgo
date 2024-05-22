package controllers

import (
	"fmt"
	"golang-for-women/sesi-6/config"
	"golang-for-women/sesi-6/models"
	"time"
)

func CreateVariant(data models.Variant) {
	var variant = models.Variant{}

	// You can further process the request body data as needed
	currentTime := time.Now()

	sqlStatement := `INSERT INTO variants (name, quantity, product_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	result, err := config.DB.Exec(sqlStatement, data.Name, data.Quantity, data.ProductID, currentTime, currentTime)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Retrieve the inserted row using the lastInsertID
	sqlRetrieve := `SELECT * FROM variants WHERE id = ?`

	err = config.DB.QueryRow(sqlRetrieve, lastInsertID).Scan(&variant.ID, &variant.Name, &variant.Quantity, &variant.ProductID, &variant.CreatedAt, &variant.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data created: %+v\n", variant)

}

func UpdateVariantById(variantID int, data models.Variant) {
	var variant = models.Variant{}

	// You can further process the request body data as needed
	currentTime := time.Now()

	sqlStatement := `UPDATE variants SET name=?, quantity=?, product_id=?, updated_at=? WHERE id=?;`

	result, err := config.DB.Exec(sqlStatement, data.Name, data.Quantity, data.ProductID, currentTime, variantID)
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
	sqlRetrieve := `SELECT * FROM variants WHERE id=?`

	err = config.DB.QueryRow(sqlRetrieve, variantID).Scan(&variant.ID, &variant.Name, &variant.Quantity, &variant.ProductID, &variant.CreatedAt, &variant.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data updated: %+v\n", variant)
}

func DeleteVariantById(variantID int) {

	sqlStatement := `DELETE FROM variants WHERE id=?;`

	result, err := config.DB.Exec(sqlStatement, variantID)
	if err != nil {
		panic(err)
	}

	// Get affected rows
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Data successfully deleted. Rows affected:", count, "row(s)")

}
