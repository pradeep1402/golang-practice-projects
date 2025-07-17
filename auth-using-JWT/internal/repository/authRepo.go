package repo

import (
	"context"
	"fmt"
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
