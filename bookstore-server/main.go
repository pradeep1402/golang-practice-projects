package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type server struct{}

func (app *server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	books, err := os.ReadFile("./books.json")

	if err != nil {
		res.Header().Set("Content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(`{"err": "404"}`))
		return
	}

	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(books))
}

type Book struct {
	Id     int
	Title  string
	Author string
	Price  float64
}

func searchBooks(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query().Get("query")
	data, err := os.ReadFile("./books.json")
	res.Header().Set("Content-type", "application-json")

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(`{"err": "404"}`))
		return
	}

	var books []Book
	err = json.Unmarshal(data, &books)

	if err != nil {
		fmt.Println("Unable to decode... Error: ", err)
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(`{"err": "Unable to decode"}`))
		return
	}

	var filteredBooks []Book

	for _, book := range books {
		if strings.Contains(book.Title, params) {
			filteredBooks = append(filteredBooks, book)
		}
	}

	fmt.Println(filteredBooks)
	bytes, err := json.Marshal(filteredBooks)

	if err != nil {
		fmt.Println("Unable to encode... Error: ", err)
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(`{"err": "Unable to encode"}`))
		return
	}

	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(bytes))
}

func main() {
	app := &server{}

	http.Handle("/", app)
	http.HandleFunc("/search", searchBooks)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
