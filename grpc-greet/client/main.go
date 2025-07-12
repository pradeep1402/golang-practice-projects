package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-service-greet/proto"
)

const addr string = "localhost:50051"

func main() {
	// connection with the server
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials((insecure.NewCredentials())))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	// greeter client
	c := pb.NewGreeterClient(conn)
	res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "pradeep"})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Greeting: %s", res.GetMessage())

	// calculator client
	calculatorClient := pb.NewCalculatorServiceClient(conn)
	sumRes, err := calculatorClient.Sum(context.Background(), &pb.OperandsRequest{Operad1: 10, Operad2: 3})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Sum Result: %d\n", sumRes.GetSum())

	// saysHello client
	result, err := c.SaysHello(context.Background(), &pb.HelloRequest{Name: "Its me..."})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	for {
		msg, err := result.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err.Error())
		}

		log.Println(msg)
	}

	// prime factor
	primeFactor, err := calculatorClient.PrimeFactor(context.Background(), &pb.Number{Number: 21})
	if err != nil {
		log.Fatalf("Error: %v\n", err.Error())
	}

	for {
		res, err := primeFactor.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err.Error())
		}

		log.Print(res.Sum)
	}

	// LongGreet
	reqs := []*pb.HelloRequest{
		{Name: "pradeep"},
		{Name: "atul"},
		{Name: "danish"},
		{Name: "charan"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Errror: %v\n", err.Error())
		return
	}

	for _, req := range reqs {
		log.Printf("Sending req for: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err = stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while sending req: %v\n", err.Error())
	}

	log.Println(res)

	// Average of list of numbers
	listOfNum := []*pb.Number{
		{Number: 20}, {Number: 9}, {Number: 40},
	}

	avgStream, err := calculatorClient.Average(context.Background())
	if err != nil {
		log.Fatalf("Error: %v\n", err.Error())
		return
	}

	for _, req := range listOfNum {
		avgStream.Send(req)
		time.Sleep(1 * time.Second)
	}

	avg, err := avgStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while sending reqs: %v\n", err.Error())
	}

	log.Println(avg)

	// Greet Everyone
	streamToGreetEveryone, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error: %v\n", err.Error())
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			streamToGreetEveryone.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			reply, err := streamToGreetEveryone.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading the stream: %v\n", err.Error())
			}
			log.Println(reply.Message)
		}
		close(waitc)
	}()
	<-waitc
}
