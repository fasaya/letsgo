package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB  *sql.DB
	err error
)

func InitDB() (*sql.DB, error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// host := os.Getenv("HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	// dbport := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	mysqlInfo := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, db_name)

	DB, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		return nil, err
	}

	return DB, nil

	// defer DB.Close()

	// err = DB.Ping()
	// if err != nil {
	// 	panic(err)
	// }
}
