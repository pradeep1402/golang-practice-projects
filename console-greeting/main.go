package main

import (
	"fmt"
	"os"
)

func greet(name string) string {
	if name == "" {
		return "Usage: go run main.go [name]"
	}

	return "Hello, " + name + "! Welcome to Go"
}

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(greet(name))
}
