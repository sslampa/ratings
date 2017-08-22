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
	port := flag.String("port", "8080", "Port to serve on")
	flag.Parse()

	models.Initialize("ratings_app")

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Printf("Serving on HTTP port: %s\n", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
