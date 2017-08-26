package handlers

import (
	"net/http"
)

// UserRoute routes to correct handler
func UserRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		PostUserHandler(w, r)
		return
	case "DELETE":
		DeleteUserHandler(w, r)
		return
	}

	// re := regexp.MustCompile("^.*/users/$")
	// if re.MatchString(r.URL.Path) {
	//   return UsersHandler
	// }

	UsersHandler(w, r)
	return
}
