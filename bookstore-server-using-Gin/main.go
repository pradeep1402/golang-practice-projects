package main

import (
	"encoding/json"
	"fmt"
	"gin/app"
	"gin/models"
	"os"
)

func main() {
	data, err := os.ReadFile("./books.json")
	var books []models.Book
	if err != nil {
		fmt.Println("Error while reading the data...")
	}

	json.Unmarshal(data, &books)

	r := app.SetupRouter(books)
	r.Run(":8080")
}
