package main

import (
	"context"
	"gin/app"
	"gin/db"
	"gin/handlers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	dbPool, err := db.ConnectDB(ctx, dbUrl)
	if err != nil {
		log.Fatal("Error while connecting to db")
	}
	defer dbPool.Close()
	bookHandler := handlers.CreateBookHandler(dbPool)

	server := app.SetupRouter(bookHandler)
	server.Run(":8080")
}
