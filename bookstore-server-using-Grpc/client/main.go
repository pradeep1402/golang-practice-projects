package main

import (
	"context"
	"log"

	pb "grpc-bookStore/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

const url = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect with the server: %v\n", err.Error())
	}
	defer conn.Close()

	bookStoreClient := pb.NewBookStoreClient(conn)

	res, err := bookStoreClient.GetById(context.Background(), &pb.BookIdRequest{Id: 1})

	if err != nil {
		log.Fatalf("Error: %v\n", err.Error())
	}

	log.Println(res)

	books, err := bookStoreClient.GetBooks(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error: %v\n", err.Error())
	}

	for _, val := range books.Books {
		log.Println(val)
	}
}
