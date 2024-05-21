package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-for-women/sesi-6/models"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
	err error
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   bool   `json:"error"`
}

var PORT = ":9090"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// host := os.Getenv("HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	// dbport := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	mysqlInfo := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, db_name)

	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Successfully connected to database", db_name)

	http.HandleFunc("/", index)
	http.HandleFunc("/products", getProductWithVariant)
	http.HandleFunc("/product/", updateProduct)
	http.HandleFunc("/product", createProduct)

	fmt.Println("Application is listening on port", PORT, "and successfully connected to database", db_name)
	http.ListenAndServe(PORT, nil)

	// createProduct()
	// updateProduct()
	// getProductById()
	// createVariant()
	// updateVariantById()
	// deleteVariantById()
	// getProductWithVariant()
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		data := map[string]string{
			"message": "Hello world!",
		}

		json.NewEncoder(w).Encode(data)
	}
}

func getProductWithVariant(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		var results = []models.Product{}

		sqlStatement := `SELECT * from products`

		rows, err := db.Query(sqlStatement)
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

			results = append(results, product)
		}

		// fmt.Println("Book datas:", results)
		json.NewEncoder(w).Encode(results)
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		w.Header().Set("Content-Type", "application/json")
		var product = models.Product{}

		// Decode the JSON body into the struct
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, "Unable to parse JSON body", http.StatusBadRequest)
			return
		}

		// You can further process the request body data as needed
		currentTime := time.Now()

		sqlStatement := `INSERT INTO products (name, created_at, updated_at) VALUES (?, ?, ?)`

		result, err := db.Exec(sqlStatement, product.Name, currentTime, currentTime)
		if err != nil {
			panic(err)
		}

		lastInsertID, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		// Retrieve the inserted row using the lastInsertID
		sqlRetrieve := `SELECT * FROM products WHERE id = ?`

		err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Data created: %+v\n", product)

		results := Response{
			Error:   false,
			Message: "Success creating product",
			Data:    product,
		}

		json.NewEncoder(w).Encode(results)
	}
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		w.Header().Set("Content-Type", "application/json")
		var product = models.Product{}

		// Split the URL path by '/' to get the individual path segments
		pathParts := strings.Split(r.URL.Path, "/")

		// The last segment should be the product ID
		id := pathParts[len(pathParts)-1]
		// fmt.Println("ID:", id)

		// check if id empty
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		// Decode the JSON body into the struct
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, "Unable to parse JSON body", http.StatusBadRequest)
			return
		}

		// You can further process the request body data as needed
		currentTime := time.Now()

		sqlStatement := `UPDATE products SET name = ?, updated_at = ? WHERE id = ?;`

		result, err := db.Exec(sqlStatement, product.Name, currentTime, id)
		if err != nil {
			panic(err)
		}

		count, err := result.RowsAffected()
		if err != nil {
			panic(err)
		}
		fmt.Println("Rows affected:", count, "row(s)")

		// Retrieve the inserted row using the id
		sqlRetrieve := `SELECT * FROM products WHERE id = ?`

		err = db.QueryRow(sqlRetrieve, id).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			w.WriteHeader(404)
			results := Response{
				Error:   true,
				Message: "ID not found",
			}
			json.NewEncoder(w).Encode(results)
			return
			// panic(err)
		}

		// fmt.Printf("Data updated: %+v\n", product)

		results := Response{
			Error:   false,
			Message: "Success update product",
			Data:    product,
		}

		json.NewEncoder(w).Encode(results)
	}
}
