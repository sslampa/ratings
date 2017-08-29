package handlers

import (
	"net/http"
)

// UserRoute routes to correct handler
func UserRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		DeleteUserHandler(w, r)
		return
	case "POST":
		PostUserHandler(w, r)
		return
	case "GET":
		UserHandler(w, r)
	}

	return
}
