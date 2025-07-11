package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-service-greet/proto"
)

const addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials((insecure.NewCredentials())))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "pradeep"})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	calculatorClient := pb.NewCalculatorServiceClient(conn)

	sumRes, err := calculatorClient.Sum(context.Background(), &pb.OperandsRequest{Operad1: 10, Operad2: 3})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Sum Result: %d\n", sumRes.GetSum())

	log.Println("All RPCs completed successfully!")
}
