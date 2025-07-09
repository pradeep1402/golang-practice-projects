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

# 🐘 Running PostgreSQL in Docker (with OrbStack)

This guide helps you run PostgreSQL inside a Docker container using OrbStack (or Docker CLI). No need to install PostgreSQL on your machine directly.

---

## 📦 What You'll Set Up

- A PostgreSQL 15 database
- Exposed on `localhost:5432`
- Default user: `postgres`
- Default password: `secret`
- Default database: `bookstore`
- Data stored persistently using Docker volumes

---

## 📄 Step 1: Create `docker-compose.yml`

Create a file named `docker-compose.yml` in your project root:

```yaml
version: "3.8"

services:
  postgres:
    image: postgres:15
    container_name: bookstore-postgres
    restart: always
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: bookstore
    ports:
      - "5432:5432"
    volumes:
      - pg_data:/var/lib/postgresql/data

volumes:
  pg_data:
```

## 🚀 Step 2: Start PostgreSQL

In your terminal, run this inside the same directory as the docker-compose.yml file:

bash
Copy
Edit

- With OrbStack (if supported)

```
orbstack compose up -d
```

- OR with Docker CLI (works in all cases)

```
docker compose up -d
```
