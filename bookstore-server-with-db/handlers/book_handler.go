package handlers

import (
	"gin/models"
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

func (s *BookHandler) GetBooks(ctx *gin.Context) {
	books, err := s.service.GetBooks(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (s *BookHandler) AddBook(ctx *gin.Context) {
	title := ctx.PostForm("title")
	author := ctx.PostForm("author")
	price, err := strconv.ParseFloat(ctx.PostForm("price"), 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	book := models.PostFormBook{Title: title, Price: price, Author: author}
	books, err := s.service.AddBook(ctx, book)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (s *BookHandler) UpdateBookPrice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
	}

	price, err := strconv.ParseFloat(ctx.PostForm("price"), 64)

	if err != nil {
		ctx.JSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
	}

	book, err := s.service.UpdateBookPrice(ctx, id, price)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (s *BookHandler) DeleteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}

	book, err := s.service.DeleteById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}
