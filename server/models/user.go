package models

import (
	"database/sql"
	"log"
)

// User fields
type User struct {
	ID       string
	Username string
}

// GetUser returns the user
func GetUser(c, v string, db *sql.DB) (User, error) {
	u := User{}
	var uQuery string
	switch c {
	case "username":
		uQuery = `SELECT * FROM users WHERE username = $1`
	case "id":
		uQuery = `SELECT * FROM users WHERE id = $1`
	}

	rows, err := db.Query(uQuery, v)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Username)
		if err != nil {
			return u, err
		}
	}

	return u, nil
}

// PostUser returns the user
func PostUser(un string, db *sql.DB) (User, error) {
	u := User{}
	postQuery := `INSERT INTO users (username) VALUES ($1)`
	rows, err := db.Query(postQuery, un)
	if err != nil {
		return u, err
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Username)
		if err != nil {
			return u, err
		}
	}

	return u, nil
}
