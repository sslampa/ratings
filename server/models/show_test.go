package models

import (
	"log"
	"testing"
)

func init() {
	Initialize("ratings_app_test")
	Seed()
}

func TestGetShows(t *testing.T) {
	shows, err := GetShows("")
	if err != nil {
		log.Fatal(err)
	}

	if len(shows) != 4 {
		t.Error("Expected length of shows to be 4")
	}
}

func TestGetShowsID(t *testing.T) {
	shows, err := GetShows("1")
	if err != nil {
		log.Fatal(err)
	}

	if len(shows) != 2 {
		t.Error("Expected length of shows to be 2")
	}

	shows, err = GetShows("1000000")
	if err != nil {
		log.Fatal(err)
	}

	if len(shows) > 0 {
		t.Error("Expected length of shows to be 0")
	}
}
