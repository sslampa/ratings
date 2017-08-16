package models

import (
	"database/sql"
	"log"
	"testing"
)

var db *sql.DB

func init() {
	db = Initialize("ratings_app_test")
}

func TestUserID(t *testing.T) {
	// db := Initialize("ratings_app_test")

	createUserTable(db)
	uExpect := User{"1", "sslampa"}
	user, err := GetUser("id", "1", db)
	if err != nil {
		log.Fatal(err)
	}

	if uExpect.ID != user.ID {
		t.Errorf("User ID %v does not match expected ID %v", user.ID, uExpect.ID)
	}

	if uExpect.Username != user.Username {
		t.Errorf("Username %v does not match expected username %v", user.Username, uExpect.Username)
	}
}

func TestUserUsername(t *testing.T) {
	createUserTable(db)
	uExpect := User{"1", "sslampa"}
	user, err := GetUser("username", "sslampa", db)
	if err != nil {
		log.Fatal(err)
	}

	if uExpect.ID != user.ID {
		t.Errorf("User ID %v does not match expected ID %v", user.ID, uExpect.ID)
	}

	if uExpect.Username != user.Username {
		t.Errorf("Username %v does not match expected username %v", user.Username, uExpect.Username)
	}
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
