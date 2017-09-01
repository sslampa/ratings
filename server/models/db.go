package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// For testing purposes
	_ "github.com/lib/pq"
)

var db *sql.DB

// Initialize initializes database for REST API
func Initialize(dbName string) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s", dbUser, dbPass, dbName)

	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("You connected to your database")
	Create()
}

// Seed seeds all tables
func Seed() {
	dropUserTable()
	dropShowsTable()
	createUserTable()
	createShowsTable()
	seedUsers()
}

// Create creates all tables
func Create() {
	createUserTable()
	createShowsTable()
}

// Drop drops all tables
func Drop() {
	dropUserTable()
	dropShowsTable()
}
