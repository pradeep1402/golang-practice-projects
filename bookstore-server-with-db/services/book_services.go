package services

import (
	"context"
	"gin/models"
)

type BookRepository interface {
	GetByID(ctx context.Context, id int) (models.Book, error)
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
