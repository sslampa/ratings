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

// SeedUsers seeds users table
func SeedUsers(db *sql.DB) {
	u1 := User{Username: "sslampa"}
	u2 := User{Username: "tomanistor"}
	u3 := User{Username: "suzmas"}
	users := []User{u1, u2, u3}

	for _, u := range users {
		_, err := PostUser(u.Username, db)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("User seed created")
}
