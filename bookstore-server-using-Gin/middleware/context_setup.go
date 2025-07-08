package middleware

import (
	"fmt"
	"gin/models"

	"github.com/gin-gonic/gin"
)

func SetContext(books []models.Book) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("books", books)
		fmt.Print("heelo")
		ctx.Next()
	}
}
