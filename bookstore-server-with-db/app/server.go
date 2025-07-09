package app

import (
	"gin/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(pool *pgxpool.Pool) *gin.Engine {
	handler := handlers.CreateBookHandler(pool)
	router := gin.Default()
	books := router.Group("/books")
	books.GET("/:id", handler.GetById)

	return router
}
