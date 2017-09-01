package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sslampa/ratings/server/handlers"
	"github.com/sslampa/ratings/server/models"
)

func main() {
	port := flags()

	r := mux.NewRouter()
	r.HandleFunc("/users/{username}/shows", handlers.UserShowsHandler)
	r.HandleFunc("/users/{username}", handlers.UserHandler).Methods("GET")
	r.HandleFunc("/users/{username}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{username}", handlers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users", handlers.PostUserHandler).Methods("POST")
	r.HandleFunc("/users", handlers.UsersHandler).Methods("GET")

	log.Printf("Serving on HTTP port: %s\n", *port)
	err := http.ListenAndServe(":"+*port, r)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}

func flags() *string {
	port := flag.String("port", "8080", "Port to serve on")
	seed := flag.Bool("seed", false, "Seed database")
	drop := flag.Bool("drop", false, "Drop database")
	flag.Parse()

	models.Initialize("ratings_app")

	if *seed {
		models.Seed()
	}

	if *drop {
		models.Drop()
	}

	return port
}
