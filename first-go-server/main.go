package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	staticFile, err := os.ReadFile("./public/index.html")

	if err != nil {
		fmt.Println("Unable to read the file. Error: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "404 not found"}`))
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(staticFile)))
}

func main() {
	s := &server{}
	log.Println("Server Started at http://localhost:8080")
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
