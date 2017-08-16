package models

import (
	"database/sql"
	"log"
	"testing"
)

func UserTest(t *testing.T) {
	db := Initialize("ratings_app_test")

	createUserTable(db)
}

func createUserTable(db *sql.DB) {
	const tableQuery = `CREATE TABLE IF NOT EXISTS users
  (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE
  )`

	if _, err := db.Exec(tableQuery); err != nil {
		log.Fatal(err)
	}
}
