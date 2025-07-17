package handlers

import (
	"context"
	pb "grpc-auth-jwt/gen"
	"grpc-auth-jwt/internal/services"
	"log"
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
