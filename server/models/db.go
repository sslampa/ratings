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
}

// SeedUsers seeds users table
func SeedUsers() {
	u1 := User{Username: "sslampa"}
	u2 := User{Username: "tomanistor"}
	u3 := User{Username: "suzmas"}
	users := []User{u1, u2, u3}

	for _, u := range users {
		_, err := PostUser(u.Username)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("User seed created")
}

// CreateUserTable creates user table
func CreateUserTable() {
	const dropQuery = `DROP TABLE users`
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatal(err)
	}
	fmt.Println("User table dropped")

	const tableQuery = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE)`

	if _, err := db.Exec(tableQuery); err != nil {
		log.Fatal(err)
	}
	fmt.Println("User table created")
}
