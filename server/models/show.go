package models

import (
	"database/sql"
	"log"
)

// Show contains detail of tv shows
type Show struct {
	ID          string
	UserID      string
	Name        string
	Year        string
	Description string
}

// GetShows returns all shows
// TODO: Add more/better validation for id parameter
func GetShows(id string) ([]Show, error) {
	var shows = make([]Show, 0)
	var query string
	var err error
	var rows *sql.Rows

	if id == "" {
		query = `SELECT * FROM shows`
		rows, err = db.Query(query)
	} else {
		query = `SELECT * FROM shows WHERE user_id = $1`
		rows, err = db.Query(query, id)
	}

	if err != nil {
		return shows, err
	}

	for rows.Next() {
		show := Show{}
		err = rows.Scan(&show.ID, &show.UserID, &show.Name, &show.Year, &show.Description)
		if err != nil {
			return shows, err
		}
		shows = append(shows, show)
	}
	return shows, nil
}

func createShowsTable() {
	query := `CREATE TABLE IF NOT EXISTS shows (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users ON DELETE CASCADE,
    name VARCHAR,
    year INTEGER,
    description text
  )`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func dropShowsTable() {
	query := `DROP TABLE IF EXISTS shows`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func seedShows() {
	query := `INSERT INTO shows (user_id, name, year, description) VALUES ($1, $2, $3, $4)`
	shows := []Show{
		{UserID: "1", Name: "Community", Year: "2009", Description: "A funny show at a community college."},
		{UserID: "1", Name: "30 Rock", Year: "2007", Description: "A funny show about the crew in sketch comedy."},
		{UserID: "2", Name: "The Office", Year: "2006", Description: "A funny show about a mundane office."},
		{UserID: "3", Name: "Parks and Recreation", Year: "2009", Description: "A funny show about a parks department"},
	}

	for _, s := range shows {
		_, err := db.Exec(query, s.UserID, s.Name, s.Year, s.Description)
		if err != nil {
			log.Fatal(err)
		}
	}
}
