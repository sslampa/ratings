package models

import (
	"fmt"
	"log"
)

// Show contains detail of tv shows
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
	query := `INSERT INTO shows (name, year, description) VALUES ($1, $2, $3)`
	shows := []Show{
		{Name: "Community", Year: "2009", Description: "A funny show at a community college."},
		{Name: "30 Rock", Year: "2007", Description: "A funny show about the crew in sketch comedy."},
		{Name: "The Office", Year: "2006", Description: "A funny show about a mundane office."},
		{Name: "Parks and Recreation", Year: "2009", Description: "A funny show about a parks department"},
	}

	for _, s := range shows {
		_, err := db.Exec(query, s.Name, s.Year, s.Description)
		if err != nil {
			log.Fatal(err)
		}
	}
}
