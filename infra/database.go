package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() *sql.DB {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Fetch environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Initialize the DB connection pool
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to initialize database connection: ", err)
	}

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	} else {
		fmt.Println("Successfully connected to the database!")
	}

	return db
}
