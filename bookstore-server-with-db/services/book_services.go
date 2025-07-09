package services

import (
	"context"
	"gin/models"
)

type BookRepository interface {
	GetByID(ctx context.Context, id int) (models.Book, error)
	GetBooks(ctx context.Context) ([]models.Book, error)
	AddBook(ctx context.Context, book models.PostFormBook) (models.Book, error)
}

type BookService struct {
	repo BookRepository
}

func CreateServices(repo BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (service *BookService) GetBookById(ctx context.Context, id int) (models.Book, error) {
	return service.repo.GetByID(ctx, id)
}

func (service *BookService) GetBooks(ctx context.Context) ([]models.Book, error) {
	return service.repo.GetBooks(ctx)
}

func (service *BookService) AddBook(ctx context.Context, book models.PostFormBook) (models.Book, error) {
	return service.repo.AddBook(ctx, book)
}
