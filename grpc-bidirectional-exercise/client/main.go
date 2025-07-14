package main

import (
	"context"
	"io"
	"log"

	pb "grpc-bidirectional/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Unable to connect: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	numbers := []*pb.MaxRequest{
		{Num: 12},
		{Num: 0},
		{Num: 15},
		{Num: 30},
	}
	log.Println("connected...")

	stream, err := client.Max(context.Background())

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	wait := make(chan struct{})
	go func() {
		for _, req := range numbers {
			err := stream.Send(req)
			if err != nil {
				log.Printf("Unable to send the stream: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			max, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading the stream: %v\n", err.Error())
			}
			log.Println(max)
		}
		close(wait)
	}()

	<-wait

	squt, err := client.Squt(context.Background(), &pb.SqutRequest{Number: -25})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Fatalf("Error message from server: %s\n", e.Message())
			log.Fatalf("Error code is: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Fatalln("Send a negative number!")
				return
			}
		} else {
			log.Fatalf("Error: %v\n", err.Error())
		}
	}

	log.Println(squt.SquareRoot)
}
