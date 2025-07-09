package handlers

import (
	"gin/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookHandler struct {
	dbPool *pgxpool.Pool
}

func CreateBookHandler(dbPool *pgxpool.Pool) *BookHandler {
	return &BookHandler{dbPool: dbPool}
}

func (db *BookHandler) GetBookById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Book
	err = db.dbPool.QueryRow(ctx.Request.Context(), "Select * from books where id = $1", id).
		Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book: " + err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (db *BookHandler) GetBooks(ctx *gin.Context) {
	rows, err := db.dbPool.Query(ctx.Request.Context(), "Select * from books")
	if err != nil {
		log.Fatal("Error failed to read.")
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to read"})
		return
	}

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan book row: " + err.Error()})
			return
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error during row iteration: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}
