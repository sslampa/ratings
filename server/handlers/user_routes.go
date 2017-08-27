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
	case "GET":
		UserHandler(w, r)
		return
	}

	// re := regexp.MustCompile("^.*/users/$")
	// if re.MatchString(r.URL.Path) {
	//   return UsersHandler
	// }

	return
}
