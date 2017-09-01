package models

import (
	"fmt"
	"log"
)

type Show struct {
	ID          string
	Name        string
	Year        string
	Description string
}

func createShowsTable() {
	query := `CREATE TABLE IF NOT EXISTS shows (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    year INTEGER,
    description text
  )`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shows table created")
}

func dropShowsTable() {
	query := `DROP TABLE IF EXISTS shows`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shows table dropped")
}

func seedShows() {

}
