package handlers

import (
	"context"
	"grpc-bookStore/internals/services"
	pb "grpc-bookStore/proto"
	"log"
)

type BookStoreHandler struct {
	services *services.BookServices
	pb.UnimplementedBookStoreServer
}

func CreateBookStoreHandler(services *services.BookServices) *BookStoreHandler {
	return &BookStoreHandler{services: services}
}

func (s *BookStoreHandler) GetById(ctx context.Context, req *pb.BookIdRequest) (*pb.BookDetailResponse, error) {
	id := int(req.GetId())
	log.Printf("GetByID is invoked with %v.\n", id)
	book, err := s.services.GetByID(ctx, id)

	if err != nil {
		log.Fatalf("Unable to fetch from db: %v\n", err.Error())
	}

	return &pb.BookDetailResponse{Id: int64(book.Id), Titile: book.Title, Author: book.Author, Price: float32(book.Price)}, err
}
