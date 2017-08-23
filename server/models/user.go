package models

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

// User fields
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// GetUsers returns all users
func GetUsers() ([]User, error) {
	var users []User

	getQuery := "SELECT * FROM users"
	rows, err := db.Query(getQuery)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Username)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}

// GetUser returns the user
func GetUser(c, v string) (User, error) {
	u := User{}
	var getQuery string
	switch c {
	case "username":
		getQuery = "SELECT * FROM users WHERE username = $1"
	case "id":
		getQuery = "SELECT * FROM users WHERE id = $1"
	default:
		return u, errors.New("Entered incorrect value for query case")
	}

	err := db.QueryRow(getQuery, v).Scan(&u.ID, &u.Username)
	if err != nil {
		return u, fmt.Errorf("No user found with %v %v", c, v)
	}

	return u, nil
}

// PostUser returns the user
func PostUser(un string) (User, error) {
	u := User{}
	id := 0

	postQuery := "INSERT INTO users (username) VALUES ($1) RETURNING id"
	err := db.QueryRow(postQuery, un).Scan(&id)
	if err != nil {
		return u, err
	}

	u, err = GetUser("id", strconv.Itoa(int(id)))
	if err != nil {
		return u, err
	}

	return u, nil
}

func seedUsers() {
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

func createUserTable() {
	const tableQuery = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE)`

	if _, err := db.Exec(tableQuery); err != nil {
		log.Fatal(err)
	}
	fmt.Println("User table created")
}

func dropUserTable() {
	const dropQuery = `DROP TABLE IF EXISTS users`
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatal(err)
	}
	fmt.Println("User table dropped")
}
