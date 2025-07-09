package app

import (
	"gin/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handlers *handlers.BookHandler) *gin.Engine {
	router := gin.Default()
	// router.GET("/", handlers.BookHandler)
	// router.POST("/", handlers.HandleAddBook)
	router.GET("book/:id", handlers.GetBookById)
	// router.DELETE("book/:id", handlers.HandleDeleteBook)

	return router
}
