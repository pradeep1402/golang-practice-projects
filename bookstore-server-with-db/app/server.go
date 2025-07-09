package app

import (
	"gin/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handlers *handlers.BookHandler) *gin.Engine {
	router := gin.Default()
	books := router.Group("/books")
	books.GET("/", handlers.GetBooks)
	// books.POST("/", handlers.HandleAddBook)
	books.GET("/:id", handlers.GetBookById)
	// books.DELETE("/:id", handlers.HandleDeleteBook)

	return router
}
