package main

import (
	"context"
	"io"
	"log"
	"net"

	pb "grpc-service-greet/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.UnimplementedGreeterServer
	pb.UnimplementedCalculatorServiceServer
}

func (s *Server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *Server) Sum(_ context.Context, in *pb.OperandsRequest) (*pb.OperandsResponse, error) {
	log.Printf("Received: %v, %v\n", in.Operad1, in.Operad2)
	return &pb.OperandsResponse{Sum: in.GetOperad1() + in.GetOperad2()}, nil
}

func (s *Server) SaysHello(in *pb.HelloRequest, stream pb.Greeter_SaysHelloServer) error {
	log.Printf("Received: %v", in.GetName())

	for range 5 {
		res := &pb.HelloReply{Message: "Hello " + in.GetName()}

		if err := stream.Send(res); err != nil {
			log.Printf("Error sending response in SaysHello: %v\n", err.Error())
			return err
		}
	}
	return nil
}

func (s *Server) PrimeFactor(in *pb.Number, stream pb.CalculatorService_PrimeFactorServer) error {
	log.Printf("Received: %v\n", in.Number)
	factor := 2
	num := in.Number
	for num > 1 {
		if num%int64(factor) == 0 {
			num = num / int64(factor)
			res := &pb.OperandsResponse{Sum: int64(factor)}
			if err := stream.Send(res); err != nil {
				log.Printf("Error sending response: %v\n", err.Error())
				return err
			}
		} else {
			factor++
		}
	}
	return nil
}

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.HelloRequest, pb.HelloReply]) error {
	log.Println("Long Greet was invoked.")
	res := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: res})
		}

		if err != nil {
			log.Fatalf("Error while reading stream in server: %v\n", err.Error())
			return err
		}

		res = res + " " + req.GetName()
	}
}

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average invoked.")

	var sum int64
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			avg := float32(sum) / float32(count)
			return stream.SendAndClose(&pb.AverageResult{Avg: avg})
		}

		if err != nil {
			log.Fatalf("Error while reading stream in server: %v\n", err.Error())
			return err
		}
		sum += req.Number
		count++
	}
}

func (s *Server) GreetEveryone(stream pb.Greeter_GreetEveryoneServer) error {
	log.Println("GreetEveryone invoked.")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err.Error())
		}

		err = stream.Send(&pb.HelloReply{Message: "hello " + req.Name})

		if err != nil {
			log.Fatalf("Error while sending the stream: %v\n", err.Error())
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	pb.RegisterCalculatorServiceServer(s, &Server{})
	log.Printf("Listening on %s\n", addr)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
