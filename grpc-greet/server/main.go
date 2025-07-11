package main

import (
	"context"
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
