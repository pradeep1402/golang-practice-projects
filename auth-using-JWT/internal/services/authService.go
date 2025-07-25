package services

import (
	"context"
	"grpc-auth-jwt/internal/models"
	repo "grpc-auth-jwt/internal/repository"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	Register(ctx context.Context, email string, password string) error
	Login(ctx context.Context, email string) (models.User, error)
}

type AuthService struct {
	repo AuthRepository
}

func CreateServices(repo *repo.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

const key = "secret-key"

// func createToken(email string) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["authorized"] = true
// 	claims["username"] = email
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

// 	tokenString, err := token.SignedString([]byte(key))
// 	if err != nil {
// 		log.Fatalf("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}

// 	return tokenString, nil
// }

func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{"email": email, "exp": time.Now().Add(time.Hour * 24).Unix()})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Fatalf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func verifyToken(token string) error {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return err
	}

	if !t.Valid {
		return err
	}

	return nil
}

func (r *AuthService) Register(ctx context.Context, email string, password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatalf("Unable to hash password: %s\n", err.Error())
		return "", err
	}

	err = r.repo.Register(ctx, email, string(hashPass))
	if err != nil {
		log.Fatalf("Unable to insert: %s\n", err.Error())
		return "", err
	}

	jwt, err := createToken(email)
	if err != nil {
		log.Fatalf("Unable to create token: %s\n", err.Error())
		return "", err
	}

	return jwt, nil
}

func (r *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := r.repo.Login(ctx, email)
	if err != nil {
		log.Printf("Unable to hash password: %s\n", err.Error())
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	jwt, err := createToken(email)

	if err != nil {
		log.Fatalf("Unable to genrate the jwt: %s\n", err.Error())
		return "", err
	}

	return jwt, nil
}

func Validate(ctx context.Context, jwt string) (bool, error) {
	err := verifyToken(jwt)

	if err != nil {
		log.Printf("Unable to verify: %s\n", err.Error())
		return false, err
	}

	return true, nil
}
