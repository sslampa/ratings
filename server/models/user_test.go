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
	expect := User{"1", "sslampa"}
	user, err := GetUser("id", "1", db)
	if err != nil {
		log.Fatal(err)
	}

	userComp(t, expect, user)
}

func TestUserUsername(t *testing.T) {
	expect := User{"1", "sslampa"}
	user, err := GetUser("username", "sslampa", db)
	if err != nil {
		log.Fatal(err)
	}

	userComp(t, expect, user)
}

func TestUserAdd(t *testing.T) {
	expect := User{"4", "cmfasulo"}
	_, err := PostUser("cmfasulo", db)
	if err != nil {
		log.Fatal(err)
	}

	user, err := GetUser("username", "cmfasulo", db)
	if err != nil {
		log.Fatal(err)
	}

	userComp(t, expect, user)

	_, err = PostUser("cmfasulo", db)
	if err == nil {
		t.Errorf("Username %v should not be the same as %v", user.Username, expect.Username)
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

	SeedUsers(db)
}

func userComp(t *testing.T, expected, actual User) {
	if expected.ID != actual.ID {
		t.Errorf("User ID %v does not match expected ID %v", actual.ID, expected.ID)
	}

	if expected.Username != actual.Username {
		t.Errorf("Username %v should not match expected username %v", actual.Username, expected.Username)
	}
}
