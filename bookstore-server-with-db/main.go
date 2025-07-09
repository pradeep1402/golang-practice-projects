package main

import (
	"encoding/json"
	"fmt"
	"gin/app"
	"gin/handlers"
	"os"
)

func main() {
	data, err := os.ReadFile("./books.json")
	if err != nil {
		fmt.Println("Error while reading the data...")
	}

	json.Unmarshal(data, &handlers.Books)

	r := app.SetupRouter()
	r.Run(":8080")
}
