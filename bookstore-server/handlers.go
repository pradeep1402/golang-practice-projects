package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func (app *server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	books, err := os.ReadFile("./books.json")

	if err != nil {
		withJson(res, 404, []byte(`{"err": "404"}`))
		return
	}

	withJson(res, 200, books)
}

func withJson(res http.ResponseWriter, statusCode int, books []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(statusCode)
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
		withJson(res, 404, []byte(`{"err": "404"}`))
		return
	}

	var books []Book
	err = json.Unmarshal(data, &books)

	if err != nil {
		fmt.Println("Unable to decode... Error: ", err)
		withJson(res, 404, []byte(`{"err": "Unable to decode"}`))
		return
	}

	var filteredBooks []Book

	for _, book := range books {
		if strings.Contains(book.Title, params) {
			filteredBooks = append(filteredBooks, book)
		}
	}

	fmt.Println(filteredBooks)
	result, err := json.Marshal(filteredBooks)

	if err != nil {
		withJson(res, 500, []byte(`{"err": "Unable to encode result"}`))
		return
	}

	withJson(res, 200, result)
}
