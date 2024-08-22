package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func Dbconn() {
	var err error

	// Replace with your own database credentials or load from environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("could not open connection: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("could not ping the database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")
}
