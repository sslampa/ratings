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
	models.Seed()
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

func TestPostUser(t *testing.T) {
	req, err := http.NewRequest("POST", "/users/add?username=melky", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostUserHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %v, instead got %v", http.StatusCreated, status)
	}

	var user models.User
	json.NewDecoder(rr.Body).Decode(&user)
	if user.Username != "melky" {
		t.Errorf("Expected %v to equal melky", user.Username)
	}
}

func TestPostUsernameEmpty(t *testing.T) {
	req, err := http.NewRequest("POST", "/users/add?username=", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostUserHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status code %v, instead got %v", http.StatusBadRequest, status)
	}
}

func TestPostUsernameSame(t *testing.T) {
	req, err := http.NewRequest("POST", "/users/add?username=sslampa", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostUserHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status code %v, instead got %v", http.StatusBadRequest, status)
	}
}

func TestDeleteUser(t *testing.T) {
	req1, err := http.NewRequest("POST", "/users/add?username=mrobock", nil)
	if err != nil {
		t.Fatal(err)
	}

	r1 := httptest.NewRecorder()
	handler := http.HandlerFunc(PostUserHandler)

	handler.ServeHTTP(r1, req1)

	req2, err := http.NewRequest("DELETE", "/users/mrobock", nil)
	if err != nil {
		t.Fatal(err)
	}

	r2 := httptest.NewRecorder()
	handler = http.HandlerFunc(DeleteUserHandler)
	handler.ServeHTTP(r2, req2)
	if status := r2.Code; status != http.StatusNoContent {
		t.Errorf("Expected status code %v, instead got %v", http.StatusNoContent, status)
	}
}
