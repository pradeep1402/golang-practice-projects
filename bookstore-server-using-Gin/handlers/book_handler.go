package handlers

import (
	"fmt"
	"gin/models"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	Books []models.Book
	mu    sync.Mutex
)

func BookHandler(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(Books)
	c.JSON(http.StatusOK, Books)
}

func genrator() func() int {
	id := 10
	return func() int {
		id++
		return id
	}
}

var idGenrator = genrator()

func HandleAddBook(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	price := c.PostForm("price")

	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusNotFound, `{"err": "price is not a float"}`)
	}

	book := models.Book{Id: idGenrator(), Title: title, Author: author, Price: floatPrice}
	Books = append(Books, book)
	c.JSON(http.StatusOK, Books)
}
