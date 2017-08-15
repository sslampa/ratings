package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Initialize initializes database for REST API
func Initialize() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=ratings_app", dbUser, dbPass)

	DB, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("You connected to your database")
}
