package main

import (
	"log"
	"net/http"
)

type server struct{}

func main() {
	app := &server{}

	http.Handle("/", app)
	http.HandleFunc("/search", searchBooks)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
