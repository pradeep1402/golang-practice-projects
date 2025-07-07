# 📚 REST API Server: In-Memory Bookstore

## 🧠 Concepts Covered

- `net/http` package for building web servers
- Routing and handling HTTP verbs (GET, POST, PUT, DELETE)
- JSON encoding/decoding using `encoding/json`
- Working with structs and slices/maps in memory
- Basic RESTful API design principles

## 🏗️ Project Description

This project is a simple REST API built using Go’s standard library. It manages a collection of books stored in memory (no database). It supports the following operations:

- `GET /books` → Get the list of all books
- `GET /books/{id}` → Get a single book by ID
- `POST /books` → Add a new book
- `PUT /books/{id}` → Update an existing book
- `DELETE /books/{id}` → Delete a book by ID

Books are represented using the following struct:

```go
type Book struct {
    ID     int     `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}
```
