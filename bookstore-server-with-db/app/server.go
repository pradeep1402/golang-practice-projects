package app

import (
	"gin/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateRouter(pool *pgxpool.Pool) *gin.Engine {
	handler := handlers.CreateBookHandler(pool)
	router := gin.Default()
	books := router.Group("/books")
	books.GET("/", handler.GetBooks)
	books.GET("/:id", handler.GetById)
	books.POST("/", handler.AddBook)

	return router
}
