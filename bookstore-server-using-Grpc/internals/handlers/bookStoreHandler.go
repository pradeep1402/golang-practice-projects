package handlers

import (
	"grpc-bookStore/internals/services"
	pb "grpc-bookStore/proto"
)

type BookStoreHandler struct {
	services *services.BookServices
	*pb.UnimplementedBookStoreServer
}

func CreateBookStoreHandler(services *services.BookServices) *BookStoreHandler {
	return &BookStoreHandler{services: services}
}
