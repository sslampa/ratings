package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// For testing purposes
	_ "github.com/lib/pq"
)

// Initialize initializes database for REST API
func Initialize(dbName string) *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s", dbUser, dbPass, dbName)

	DB, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("You connected to your database")

	return DB
}
