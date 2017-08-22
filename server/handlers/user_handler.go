package handlers

import (
	"net/http"
)

// UserHandler Returns Handler
func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
