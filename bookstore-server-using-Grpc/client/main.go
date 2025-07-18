package main

import (
	"context"
	"log"

	pb "grpc-bookStore/proto"

	pb2 "github.com/pradeep1402/golang-practice-projects/shared-proto/gen/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	url     = "localhost:50051"
	authUrl = "localhost:50052"
)

type AuthServiceClient struct {
	Client pb2.AuthClient
	Conn   *grpc.ClientConn
}

type BookStoreServiceClient struct {
	Client pb.BookStoreClient
	Conn   *grpc.ClientConn
}

func setupClients() (*AuthServiceClient, *BookStoreServiceClient) {
	authConn, err := grpc.NewClient(authUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}

	bookConn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to bookstore service: %v", err)
	}

	authClient := &AuthServiceClient{
		Client: pb2.NewAuthClient(authConn),
		Conn:   authConn,
	}

	bookClient := &BookStoreServiceClient{
		Client: pb.NewBookStoreClient(bookConn),
		Conn:   bookConn,
	}

	return authClient, bookClient
}

func createAuthContext(token string) context.Context {
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + token,
	})

	return metadata.NewOutgoingContext(context.Background(), md)
}

func getBookByID(client pb.BookStoreClient, ctx context.Context, id int) {
	book, err := client.GetById(ctx, &pb.BookIdRequest{Id: int64(id)})
	if err != nil {
		log.Fatalf("Failed to get book by ID: %v", err)
	}
	log.Println("Book by ID:", book)
}

func getAllBooks(client pb.BookStoreClient, ctx context.Context) {
	books, err := client.GetBooks(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Failed to get books: %v", err)
	}

	log.Println("All books:")
	for _, b := range books.Books {
		log.Println(b)
	}
}

func login(authClient *AuthServiceClient, req *pb2.UserDetailRequest) *pb2.JwtResponse {
	token, err := authClient.Client.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("Registration failed: %v", err)
	}
	return token
}

func registerUser(authClient *AuthServiceClient, req *pb2.UserDetailRequest) *pb2.JwtResponse {
	token, err := authClient.Client.Register(context.Background(), req)
	if err != nil {
		log.Fatalf("Registration failed: %v", err)
	}
	return token
}

func main() {
	authClient, bookStoreClient := setupClients()
	defer authClient.Conn.Close()
	defer bookStoreClient.Conn.Close()

	req := &pb2.UserDetailRequest{
		Email:    "pradeep@mail.com",
		Password: "Pradeep12@",
	}

	// token := registerUser(authClient, req)
	token := login(authClient, req)

	ctx := createAuthContext(token.Jwt)
	getBookByID(bookStoreClient.Client, ctx, 1)
	getAllBooks(bookStoreClient.Client, ctx)
}
