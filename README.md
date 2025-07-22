# Go Practice Projects Collection 📚

A collection of Go projects for learning and practicing different concepts in Go programming language.

## 🎯 Key Projects

### 1. Bookstore Server using Gin

REST API server implementing a bookstore with in-memory storage using the Gin framework.

#### Features

- CRUD operations for books
- JSON handling
- In-memory data storage
- Concurrent access handling with mutex
- Path parameters

#### API Endpoints

| Method | Path        | Description       |
| ------ | ----------- | ----------------- |
| GET    | `/`         | List all books    |
| POST   | `/`         | Add a new book    |
| GET    | `/book/:id` | Get book by ID    |
| DELETE | `/book/:id` | Delete book by ID |

#### Book Structure

```go
type Book struct {
    Id     int     `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}
```

### 2. Auth using JWT

A robust authentication system implementing JSON Web Tokens (JWT) for secure user authentication and authorization.

#### Features

- User registration and login
- JWT token generation and validation
- Protected routes
- PostgreSQL database integration
- Docker containerization

### 3. Bookstore Server with gRPC

A modern bookstore service using gRPC for efficient client-server communication.

#### Features

- Bidirectional streaming
- Protocol buffers for data serialization
- PostgreSQL database integration
- Docker support
- Client implementation included

### 4. Console Greeting

A simple CLI application demonstrating fundamental Go concepts.

#### Features

- Command-line argument handling
- Unit testing examples
- Basic Go syntax and structures
- Error handling patterns

### 5. First Go Server

A basic HTTP server implementation showcasing Go's standard library capabilities.

#### Features

- Static file serving
- Basic routing
- HTTP request/response handling
- HTML template rendering

### 6. gRPC Bidirectional Exercise

An example of bidirectional streaming in gRPC with a calculator service.

#### Features

- Streaming RPC implementation
- Protocol buffer definitions
- Error handling in streams
- Client-server communication patterns

## 🧠 Concepts Covered

- **HTTP Servers**: Using both standard library and Gin framework
- **Concurrent Programming**: Mutex for safe concurrent access
- **API Design**: RESTful principles and route handling
- **Data Structures**: Working with structs, slices, and maps
- **Error Handling**: Proper error management in Go
- **JSON**: Encoding/Decoding of JSON data
- **gRPC**: Protocol buffers and service implementation
- **Testing**: Unit tests for Go applications

## 🚀 Getting Started

1. Clone the repository
2. Navigate to specific project directory
3. Run `go mod tidy` to install dependencies
4. Follow project-specific README for further instructions

## 📝 Learning Resources

- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://github.com/gin-gonic/gin)
- [gRPC in Go](https://grpc.io/docs/languages/go/quickstart/)

## 🤝 Contributing

Feel free to:

- Fork the repository
- Create a feature branch
- Submit pull requests
