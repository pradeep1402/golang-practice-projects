package handlers

import (
	"context"
	"grpc-bookStore/internals/services"
	pb "grpc-bookStore/proto"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
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
	log.Printf("GetByID invoked with %v.\n", id)
	book, err := s.services.GetByID(ctx, id)

	if err != nil {
		log.Fatalf("Unable to fetch from db: %v\n", err.Error())
		return nil, err
	}

	return &pb.BookDetailResponse{Id: int64(book.Id), Title: book.Title, Author: book.Author, Price: float32(book.Price)}, nil
}

func (s *BookStoreHandler) GetBooks(ctx context.Context, req *emptypb.Empty) (*pb.BooksDetailResponse, error) {
	log.Println("GetBooks invoked!")
	books, err := s.services.GetBooks(ctx)
	if err != nil {
		log.Fatalf("Unable to fetch from db: %v\n", err.Error())
		return nil, err
	}

	var pbBooks []*pb.BookDetailResponse
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.BookDetailResponse{
			Id:     int64(book.Id),
			Title:  book.Title,
			Author: book.Author,
			Price:  float32(book.Price),
		})
	}

	return &pb.BooksDetailResponse{Books: pbBooks}, nil
}
