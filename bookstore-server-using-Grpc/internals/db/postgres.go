package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB(ctx context.Context) (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbUrl := os.Getenv("DATABASE_URL")
	dbPool, err := pgxpool.New(ctx, dbUrl)
	return dbPool, err
}
