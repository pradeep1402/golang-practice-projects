package main

import (
	"context"
	"grpc-bookStore/internals/db"
	"grpc-bookStore/internals/handlers"
	"grpc-bookStore/internals/repository"
	"grpc-bookStore/internals/services"
	pb "grpc-bookStore/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const addr = ("localhost:50051")

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Unable to create server: %s\n", err.Error())
	}

	dbCtx := context.Background()
	dbPool, err := db.ConnectDB(dbCtx)
	repo := repository.CreateBookStoreRepo(dbPool)
	services := services.CreateBookStoreServices(repo)
	handlers := handlers.CreateBookStoreHandler(services)

	if err != nil {
		log.Fatalf("Unable to connect with DB: %v\n", err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterBookStoreServer(server, handlers)
	log.Printf("Listening on %s\n", addr)

	if err = server.Serve(lis); err != nil {
		log.Printf("Failed to connect: %v\n", err.Error())
	}
}
