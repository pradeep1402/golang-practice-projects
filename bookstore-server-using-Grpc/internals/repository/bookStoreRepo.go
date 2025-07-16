package repository

import (
	"context"
	"grpc-bookStore/internals/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct {
	pool *pgxpool.Pool
}

func CreateBookStoreRepo(pool *pgxpool.Pool) *BookRepository {
	return &BookRepository{pool: pool}
}

func (db *BookRepository) GetByID(ctx context.Context, id int) (models.Book, error) {
	var book models.Book
	err := db.pool.QueryRow(ctx, "SELECT * from books WHERE id = $1", id).
		Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.CreatedAt, &book.UpdatedAt)

	return book, err
}

func (db *BookRepository) GetBooks(ctx context.Context) ([]models.Book, error) {
	var books []models.Book

	rows, err := db.pool.Query(ctx, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.CreatedAt, &book.UpdatedAt)
		books = append(books, book)
	}

	return books, err
}
