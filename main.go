package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sslampa/ratings/server/handlers"
	"github.com/sslampa/ratings/server/models"
)

func main() {
	port := flags()
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/users/", handlers.UserRoute)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Printf("Serving on HTTP port: %s\n", *port)
	err := http.ListenAndServe(":"+*port, nil)
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
