package main

import (
	"context"
	"log"

	pb "github.com/pradeep1402/golang-practice-projects/shared-proto/gen/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const url = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to connect with the server: %v\n", err.Error())
	}
	defer conn.Close()

	authClient := pb.NewAuthClient(conn)

	// res, err := authClient.Register(context.Background(),
	// &pb.UserDetailRequest{Email: "pradeep@mail.com", Password: "Pradeep12@"})

	// loggingResult(err, res)

	res, err := authClient.Login(context.Background(), &pb.UserDetailRequest{Email: "pradeep@mail.com", Password: "Pradeep12@"})

	loggingResult(err, res)
}

func loggingResult(err error, res *pb.JwtResponse) {
	if err != nil {
		log.Fatalf("Unable to Register: %s\n", err.Error())
	}

	log.Printf("Jwt token: %s\n", res.GetJwt())
}
