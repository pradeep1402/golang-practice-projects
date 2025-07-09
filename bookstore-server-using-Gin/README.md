# 📚 REST API Server – In-Memory Bookstore (Gin + Go)

## 🧠 Concepts Covered

- Gin framework for building HTTP APIs in Go
- RESTful route conventions (GET, POST, PUT, DELETE)
- Path parameters with `:id`
- JSON encoding/decoding
- In-memory data storage (slice or map)
- Modular Go application structure

---

## 🚀 Project Overview

This project is a simple REST API built using the [Gin](https://github.com/gin-gonic/gin) framework. It simulates a Bookstore backend with CRUD operations, all running in memory without a database.

### ✅ Implemented Routes

| Method | Path         | Description             |
| ------ | ------------ | ----------------------- |
| GET    | `/books`     | List all books          |
| POST   | `/books`     | Add a new book          |
| GET    | `/books/:id` | Get a specific book     |
| PUT    | `/books/:id` | Update an existing book |
| DELETE | `/books/:id` | Delete a book           |

---

## 📘 Example Book Struct

```go
type Book struct {
    Id     int     `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}
```
