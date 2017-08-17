package models

import (
	"database/sql"
	"log"
	"testing"
)

var db *sql.DB

func init() {
	db = Initialize("ratings_app_test")
	createUserTable(db)
}

func TestUserID(t *testing.T) {
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

func TestUserAdd(t *testing.T) {
	uExpect := User{"2", "tomanistor"}
	_, err := PostUser("tomanistor", db)
	if err != nil {
		log.Fatal(err)
	}

	user, err := GetUser("username", "tomanistor", db)
	if err != nil {
		log.Fatal(err)
	}

	if uExpect.ID != user.ID {
		t.Errorf("User ID %v does not match expected ID %v", user.ID, uExpect.ID)
	}

	if uExpect.Username != user.Username {
		t.Errorf("Username %v does not match expected username %v", user.Username, uExpect.Username)
	}

	_, err = PostUser("tomanistor", db)
	if err == nil {
		t.Errorf("Username %v should not be the same as %v", user.Username, uExpect.Username)
	}

}

func createUserTable(db *sql.DB) {
	const dropQuery = `DROP TABLE users`
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatal(err)
	}

	const tableQuery = `CREATE TABLE IF NOT EXISTS users
  (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE
  )`

	if _, err := db.Exec(tableQuery); err != nil {
		log.Fatal(err)
	}

	const seedQuery = `INSERT INTO users (username) VALUES ($1)`
	if _, err := db.Exec(seedQuery, "sslampa"); err != nil {
		log.Fatal(err)
	}
}
