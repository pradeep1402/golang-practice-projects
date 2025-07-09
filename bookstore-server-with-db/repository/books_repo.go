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
