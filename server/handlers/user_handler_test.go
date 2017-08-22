package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sslampa/ratings/server/models"
)

var db *sql.DB

func init() {
	models.Initialize("ratings_app_test")
	models.CreateUserTable()
	models.SeedUsers()
}

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UsersHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, instead got %v", http.StatusOK, status)
	}

	users := make([]models.User, 0)
	json.NewDecoder(rr.Body).Decode(&users)
	if len(users) != 3 {
		t.Errorf("Expected length to be 3")
	}
}
