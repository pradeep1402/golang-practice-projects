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

func HandleBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for _, v := range Books {
		if v.Id == int(id) {
			ctx.JSON(http.StatusOK, v)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
}

func HandleDeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var books []models.Book

	for _, v := range Books {
		if v.Id != int(id) {
			books = append(books, v)
		}
	}
	Books = books

	ctx.JSON(http.StatusOK, Books)
}
