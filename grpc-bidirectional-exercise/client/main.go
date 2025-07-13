package main

import (
	"context"
	"io"
	"log"

	pb "grpc-bidirectional/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Unable to connect: %v\n", err)
	}
	defer conn.Close()

	maxClient := pb.NewCalculatorServiceClient(conn)

	numbers := []*pb.MaxRequest{
		{Num: 12},
		{Num: 0},
		{Num: 15},
		{Num: 30},
	}
	log.Println("connected...")

	stream, err := maxClient.Max(context.Background())

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
}
