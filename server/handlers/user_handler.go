package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

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
	un := r.URL.Path
	re := regexp.MustCompile("^.*/users/([0-9a-zA-z]+)")
	str := re.FindStringSubmatch(un)

	user, err := models.GetUser("username", str[1])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
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

// DeleteUserHandler deletes user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	un := r.URL.Path
	re := regexp.MustCompile("^.*/users/([0-9a-zA-z]+)")
	str := re.FindStringSubmatch(un)

	err := models.DeleteUser("username", str[1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
