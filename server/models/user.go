package models

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

// User fields
type User struct {
	ID       string
	Username string
}

// GetUser returns the user
func GetUser(c, v string, db *sql.DB) (User, error) {
	u := User{}
	var getQuery string
	switch c {
	case "username":
		getQuery = `SELECT * FROM users WHERE username = $1`
	case "id":
		getQuery = `SELECT * FROM users WHERE id = $1`
	default:
		return u, errors.New("Entered incorrect value for query case")
	}

	rows, err := db.Query(getQuery, v)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Username)
		if err != nil {
			return u, err
		}
	}

	if (User{}) == u {
		return u, errors.New("No user found")
	}

	return u, nil
}

// PostUser returns the user
func PostUser(un string, db *sql.DB) (User, error) {
	u := User{}
	id := 0

	postQuery := "INSERT INTO users (username) VALUES ($1) RETURNING id"
	err := db.QueryRow(postQuery, un).Scan(&id)
	if err != nil {
		return u, err
	}

	u, err = GetUser("id", strconv.Itoa(int(id)), db)
	if err != nil {
		return u, err
	}

	return u, nil
}
