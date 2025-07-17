package main

import (
	"context"
	"log"
	"net"

	pb "github.com/pradeep1402/golang-practice-projects/shared-proto/gen/auth"
	"grpc-auth-jwt/internal/db"
	"grpc-auth-jwt/internal/handlers"
	repo "grpc-auth-jwt/internal/repository"
	"grpc-auth-jwt/internal/services"

	"google.golang.org/grpc"
)

const addr = "localhost:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Unable to connect: %s\n", err.Error())
	}

	ctx := context.Background()
	dbPool, err := db.ConnectDB(ctx)

	if err != nil {
		log.Fatalf("Unable to connect with db: %s\n", err)
	}
	log.Println("Successfully connected!!!")

	repo := repo.CreateRepository(dbPool)
	services := services.CreateServices(repo)
	handler := handlers.CreateHandlers(services)

	server := grpc.NewServer()
	pb.RegisterAuthServer(server, handler)

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Unable to serve: %s\n", err.Error())
	}
}
