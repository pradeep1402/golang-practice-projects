package handlers

import (
	"gin/repository"
	"gin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookHandler struct {
	service *services.BookService
}

func CreateBookHandler(pool *pgxpool.Pool) *BookHandler {
	repo := repository.CreateRepository(pool)
	service := services.CreateServices(repo)
	return &BookHandler{service: service}
}

func (s *BookHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	book, err := s.service.GetBookById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}
