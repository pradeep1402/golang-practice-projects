package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func bookHandler(c *gin.Context) {
	books := c.MustGet("books")
	fmt.Println(books)
	c.JSON(http.StatusOK, books)
}

func setContext(books []Book) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("books", books)
		fmt.Print("heelo")
		ctx.Next()
	}
}

func setupRouter(books []Book) *gin.Engine {
	r := gin.Default()

	r.Use(setContext(books))

	r.GET("/", bookHandler)

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	return r
}

type Book struct {
	Id     int
	Title  string
	Author string
	Price  float64
}

func main() {
	data, err := os.ReadFile("./books.json")
	var books []Book
	if err != nil {
		fmt.Println("Error while reading the data...")
	}

	json.Unmarshal(data, &books)

	r := setupRouter(books)
	r.Run(":8080")
}
