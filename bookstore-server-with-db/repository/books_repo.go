package repository

import (
	"context"
	"gin/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct {
	pool *pgxpool.Pool
}

func CreateRepository(pool *pgxpool.Pool) *BookRepository {
	return &BookRepository{pool: pool}
}

func (repo *BookRepository) GetByID(ctx context.Context, id int) (models.Book, error) {
	var book models.Book
	err := repo.pool.QueryRow(ctx, "Select * from books where id = $1", id).
		Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.CreatedAt, &book.UpdatedAt)

	return book, err
}

func (repo *BookRepository) GetBooks(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	rows, err := repo.pool.Query(ctx, "Select * from books")

	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.CreatedAt, &book.UpdatedAt)
		books = append(books, book)
	}

	return books, err
}

func (repo *BookRepository) AddBook(ctx context.Context, book models.PostFormBook) (models.Book, error) {
	var insertedBook models.Book

	err := repo.pool.QueryRow(ctx, "Insert into books (title, author, price) values($1, $2, $3)  RETURNING *", book.Title, book.Author, book.Price).Scan(&insertedBook.Id, &insertedBook.Title, &insertedBook.Author, &insertedBook.Price, &insertedBook.CreatedAt, &insertedBook.UpdatedAt)

	return insertedBook, err
}
