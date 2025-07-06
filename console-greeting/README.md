1. Hello Go & CLI Greeting App
⏱ Duration: 30 minutes – 1 hour

🧠 Concepts: Go CLI basics, package structure, variables, functions, os.Args, fmt.Println, Go workspace

🔖 Tags: basics, cli, packages, functions

📝 Description:
This is your Hello World for Go — but slightly enhanced.

You’ll build a simple command-line application that takes a name as an argument and prints a personalized greeting like:

bash
Copy
Edit
$ go run main.go Pradeep
Hello, Pradeep! Welcome to Go.
If no name is provided, it should show a default message or usage instructions:

bash
Copy
Edit
$ go run main.go
Usage: go run main.go [name]
🎯 Learning Goals:
This project will help you understand:

How Go programs are structured (main.go, package main, func main)

The Go toolchain (go run, go build)

How to access command-line arguments using os.Args

Basic input validation and branching (if, len)

Formatting output using fmt.Printf, fmt.Println

How to write reusable functions (e.g., func greet(name string) string)

How Go handles variable declaration (:= vs var)

Organizing your first Go project with Go Modules (go mod init)
