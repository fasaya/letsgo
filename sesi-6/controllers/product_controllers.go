package product_controllers

import (
	"encoding/json"
	"golang-for-women/sesi-6/config"
	"golang-for-women/sesi-6/helpers"
	"golang-for-women/sesi-6/models"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		helpers.HandleSuccessfulResponse(w, "Hello world!", nil, 200)
		return
	}
}

func GetProductWithVariant(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

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

		helpers.HandleSuccessfulResponse(w, "Success creating product", products, 200)
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		w.Header().Set("Content-Type", "application/json")
		var product = models.Product{}

		// Decode the JSON body into the struct
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			helpers.HandleErrorResponse(w, "Unable to parse JSON body", 400)
			return
		}

		// You can further process the request body data as needed
		currentTime := time.Now()

		sqlStatement := `INSERT INTO products (name, created_at, updated_at) VALUES (?, ?, ?)`

		result, err := config.DB.Exec(sqlStatement, product.Name, currentTime, currentTime)
		if err != nil {
			helpers.HandleErrorResponse(w, "Error inserting data", 404)
			return
			// panic(err)
		}

		lastInsertID, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		// Retrieve the inserted row using the lastInsertID
		sqlRetrieve := `SELECT * FROM products WHERE id = ?`

		err = config.DB.QueryRow(sqlRetrieve, lastInsertID).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			helpers.HandleErrorResponse(w, "ID not found", 404)
			return
			// panic(err)
		}

		// fmt.Printf("Data created: %+v\n", product)

		helpers.HandleSuccessfulResponse(w, "Success creating product", product, 200)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		w.Header().Set("Content-Type", "application/json")
		var product = models.Product{}

		id, err := helpers.GetIdFromURLParam(r)
		if err != nil {
			helpers.HandleErrorResponse(w, "ID is required", 400)
			return
		}

		// Decode the JSON body into the struct
		err = json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			panic(err)
			// helpers.HandleErrorResponse(w, "Unable to parse JSON body", 400)
			// return
		}

		// You can further process the request body data as needed
		currentTime := time.Now()

		sqlStatement := `UPDATE products SET name = ?, updated_at = ? WHERE id = ?;`

		result, err := config.DB.Exec(sqlStatement, product.Name, currentTime, id)
		if err != nil {
			panic(err)
		}

		// Get affected rows
		_, err = result.RowsAffected()
		if err != nil {
			// helpers.HandleErrorResponse(w, "ID not found", 404)
			// return

			panic(err)
		}
		// fmt.Println("Rows affected:", count, "row(s)")

		// Retrieve the inserted row using the id
		sqlRetrieve := `SELECT * FROM products WHERE id = ?`

		err = config.DB.QueryRow(sqlRetrieve, id).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			// helpers.HandleErrorResponse(w, "ID not found", 404)
			// return
			panic(err)
		}

		helpers.HandleSuccessfulResponse(w, "Success update product", product, 200)
		return
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		var product = models.Product{}

		id, err := helpers.GetIdFromURLParam(r)
		if err != nil {
			helpers.HandleErrorResponse(w, "ID is required", 400)
			return
		}

		// Retrieve the inserted row using the id
		sqlRetrieve := `SELECT * FROM products WHERE id = ?`

		err = config.DB.QueryRow(sqlRetrieve, id).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			helpers.HandleErrorResponse(w, "ID not found", 404)
			return
			// panic(err)
		}

		helpers.HandleSuccessfulResponse(w, "Success get product", product, 200)
	}
}
