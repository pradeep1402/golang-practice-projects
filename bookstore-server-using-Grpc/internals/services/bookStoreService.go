package services

import "grpc-bookStore/internals/repository"

type BookRepository interface {
}

type BookServices struct {
	repo BookRepository
}

func CreateBookStoreServices(repo *repository.BookRepository) *BookServices {
	return &BookServices{repo: repo}
}
