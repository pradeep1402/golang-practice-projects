package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	log.Println("Successfully connected to PostgreSQL database using pgxpool!")
	return pool, err
}
