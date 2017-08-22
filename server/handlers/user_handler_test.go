package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sslampa/ratings/server/models"
)

var db *sql.DB

func init() {
	db = models.Initialize("ratings_app_test")
	models.CreateUserTable(db)
	models.SeedUsers(db)
}

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, instead got %v", http.StatusOK, status)
	}

	body := rr.Body.String()
	fmt.Println("This", body)
}
