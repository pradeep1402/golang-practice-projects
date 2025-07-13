package main

import (
	pb "grpc-bidirectional/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCalculatorServiceServer
}

const add = ("localhost:50051")

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Invoked Max function...")

	maxNum := 0
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err.Error())
		}

		maxNum = max(int(res.Num), maxNum)
		err = stream.Send(&pb.MaxResponse{Max: int64(maxNum)})

		if err != nil {
			log.Fatalf("Error: %v\n", err.Error())
		}
	}

}

func main() {
	lis, err := net.Listen("tcp", add)

	if err != nil {
		log.Fatalf("Error %v\n", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Listening to: " + add)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err.Error())
	}
}
