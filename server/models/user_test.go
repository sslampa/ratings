package models

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

var db *sql.DB

func init() {
	db = Initialize("ratings_app_test")
	createUserTable(db)
	SeedUsers(db)
}

func TestUserID(t *testing.T) {
	expect := User{"1", "sslampa"}
	user, err := GetUser("id", "1", db)
	if err != nil {
		log.Fatal(err)
	}

	userComp(t, expect, user)

	_, err = GetUser("id", "3000", db)
	if err == nil {
		t.Errorf("Expected to find no user")
	}

	_, err = GetUser("something", "1", db)
	if err == nil {
		t.Errorf("Expected to find no user")
	}
}

func TestUserUsername(t *testing.T) {
	expect := User{"1", "sslampa"}
	user, err := GetUser("username", "sslampa", db)
	if err != nil {
		log.Fatal(err)
	}

	userComp(t, expect, user)
}

func TestUserPost(t *testing.T) {
	expect := User{"4", "cmfasulo"}
	user, err := PostUser("cmfasulo", db)
	if err != nil {
		t.Errorf("Expected query to return a user")
	}

	userComp(t, expect, user)

	_, err = PostUser("cmfasulo", db)
	if err == nil {
		t.Errorf("Expected username %v to not equal %v", user.Username, expect.Username)
	}

}

func createUserTable(db *sql.DB) {
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

func userComp(t *testing.T, expected, actual User) {
	if expected.ID != actual.ID {
		t.Errorf("Expected id %v to equal %v", actual.ID, expected.ID)
	}

	if expected.Username != actual.Username {
		t.Errorf("Expected username %v to equal %v", actual.Username, expected.Username)
	}
}
