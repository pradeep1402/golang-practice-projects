package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookHandler(c *gin.Context) {
	books := c.MustGet("books")
	fmt.Println(books)
	c.JSON(http.StatusOK, books)
}
