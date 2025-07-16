package services

import (
	"context"
	"grpc-bookStore/internals/models"
	"grpc-bookStore/internals/repository"
)

type BookRepository interface {
	GetByID(ctx context.Context, id int) (models.Book, error)
	GetBooks(ctx context.Context) ([]models.Book, error)
}

type BookServices struct {
	repo BookRepository
}

func CreateBookStoreServices(repo *repository.BookRepository) *BookServices {
	return &BookServices{repo: repo}
}

func (service *BookServices) GetByID(ctx context.Context, id int) (models.Book, error) {
	return service.repo.GetByID(ctx, id)
}

func (service *BookServices) GetBooks(ctx context.Context) ([]models.Book, error) {
	return service.repo.GetBooks(ctx)
}
