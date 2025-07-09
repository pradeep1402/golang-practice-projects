package app

import (
	"gin/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", handlers.BookHandler)
	router.POST("/", handlers.HandleAddBook)
	router.GET("book/:id", handlers.HandleBook)

	return router
}
