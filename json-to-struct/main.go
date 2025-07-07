package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name     string `json:"name"`
	Age      int
	Email    string
	IsMember bool
}

func main() {
	data, err := os.ReadFile("./people.json")

	if err != nil {
		fmt.Println("Unable to read the file. Error: ", err)
		return
	}

	var people []Person
	err = json.Unmarshal(data, &people)

	if err != nil {
		fmt.Println("Unable to decode... Error: ", err)
		return
	}

	fmt.Println(people)
}
