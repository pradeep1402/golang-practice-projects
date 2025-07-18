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
	"strings"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	addr = "localhost:50051"
	key  = "secret-key"
)

func verifyToken(token string) error {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return err
	}

	if !t.Valid {
		return err
	}

	return nil
}

func authUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if info.FullMethod == "/bookStore.BookStore/Login" || info.FullMethod == "/bookStore.BookStore/Register" {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "Missing metadata")
		}
		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "Missing authorization token")
		}

		token := strings.TrimPrefix(authHeader[0], "Bearer ")
		err := verifyToken(token)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return handler(ctx, req)
	}
}

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

	server := grpc.NewServer(grpc.UnaryInterceptor(authUnaryInterceptor()))
	pb.RegisterBookStoreServer(server, handlers)
	log.Printf("Listening on %s\n", addr)

	if err = server.Serve(lis); err != nil {
		log.Printf("Failed to connect: %v\n", err.Error())
	}
}
