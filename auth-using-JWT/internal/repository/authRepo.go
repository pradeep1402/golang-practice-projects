package repo

import (
	"context"
	"fmt"
	"grpc-auth-jwt/internal/models"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	pool *pgxpool.Pool
}

func CreateRepository(pool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{pool: pool}
}

func (pool *AuthRepository) Register(ctx context.Context, email string, hashedPassword string) error {
	res, err := pool.pool.Exec(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", email, hashedPassword)
	if err != nil {
		log.Printf("Insert failed: %v", err)
		return err
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		log.Printf("No rows inserted.")
		return fmt.Errorf("no rows inserted")
	}

	log.Printf("User registered successfully: %s", email)
	return nil
}

func (pool *AuthRepository) Login(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := pool.pool.QueryRow(ctx, "SELECT email, password FROM users WHERE email = $1", email).
		Scan(&user.Email, &user.Password)

	return user, err
}
