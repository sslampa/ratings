package models

import (
	"errors"
	"fmt"
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
