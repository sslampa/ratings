package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sslampa/ratings/server/models"
)

func init() {
	models.Initialize("ratings_app_test")
	models.Seed()
}

func TestGetUserRoute(t *testing.T) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/sslampa", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UserRoute)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, instead got %v", http.StatusOK, status)
	}
}

func TestPostUserRoute(t *testing.T) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/users/?username=stephchin", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UserRoute)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %v, instead got %v", http.StatusCreated, status)
	}
}

func TestDeleteUserRoute(t *testing.T) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/users/stephchin", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UserRoute)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusNoContent {
		t.Errorf("Expected status code %v, insteaad fot %v", http.StatusNoContent, status)
	}
}
