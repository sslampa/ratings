package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sslampa/ratings/server/models"
)

// UsersHandler Returns Handler
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// PostUserHandler returns handler
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
