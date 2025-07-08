package app

import (
	"gin/handlers"
	"gin/middleware"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(books []models.Book) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.SetContext(books))

	r.GET("/", handlers.BookHandler)

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	})

	return r
}
