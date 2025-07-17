package handlers

import (
	"context"
	"grpc-auth-jwt/internal/services"
	"log"

	pb "github.com/pradeep1402/golang-practice-projects/shared-proto/gen/auth"
)

type AuthHandler struct {
	services *services.AuthService
	pb.UnimplementedAuthServer
}

func CreateHandlers(services *services.AuthService) *AuthHandler {
	return &AuthHandler{services: services}
}

func (s *AuthHandler) Register(ctx context.Context, req *pb.UserDetailRequest) (*pb.JwtResponse, error) {
	log.Println("Register invoked!")

	jwt, err := s.services.Register(ctx, req.Email, req.Password)

	if err != nil {
		return nil, err
	}

	return &pb.JwtResponse{Jwt: jwt}, nil
}

func (s *AuthHandler) Login(ctx context.Context, req *pb.UserDetailRequest) (*pb.JwtResponse, error) {
	log.Println("Login invoked!")
	res, err := s.services.Login(ctx, req.Email, req.Password)

	if err != nil {
		return nil, err
	}

	return &pb.JwtResponse{Jwt: res}, err
}
