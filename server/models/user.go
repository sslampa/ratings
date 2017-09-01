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
	Shows    []Show `json:"shows,omitempty"`
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

// GetUserShows gets the collection of shows for user
func GetUserShows(un string) (User, error) {
	u, err := GetUser("username", un)
	if err != nil {
		return u, err
	}

	shows, err := GetShows(u.ID)
	if err != nil {
		return u, err
	}

	u.Shows = shows
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

// DeleteUser deletes the given user
func DeleteUser(c, v string) error {
	var deleteQuery string
	switch c {
	case "username":
		deleteQuery = "DELETE FROM users WHERE username = $1 RETURNING id"
	case "id":
		deleteQuery = "DELETE FROM users WHERE id = $1 RETURNING id"
	default:
		return errors.New("Entered incorrect value for query case")
	}

	result, err := db.Exec(deleteQuery, v)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return fmt.Errorf("No record found with %v %v", c, v)
	}

	return nil
}

// UpdateUser updates the given user
func UpdateUser(c, v, nv string) (User, error) {
	var updateQuery string
	switch c {
	case "username":
		updateQuery = "UPDATE users SET username = $1 WHERE username = $2 RETURNING *"
	case "id":
		updateQuery = "UPDATE users SET username = $1 WHERE id = $2 RETURNING *"
	}

	u := User{}
	err := db.QueryRow(updateQuery, nv, v).Scan(&u.ID, &u.Username)
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
}

func createUserTable() {
	const tableQuery = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE)`

	if _, err := db.Exec(tableQuery); err != nil {
		log.Fatal(err)
	}
}

func dropUserTable() {
	const dropQuery = `DROP TABLE IF EXISTS users`
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatal(err)
	}
}
