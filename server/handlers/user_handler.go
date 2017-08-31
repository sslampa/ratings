package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// UserHandler returns handler
func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := models.GetUser("username", vars["username"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUserHandler deletes user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := models.DeleteUser("username", vars["username"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PostUserHandler returns handler
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	u, err := models.PostUser(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
