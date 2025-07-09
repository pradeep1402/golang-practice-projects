package main

import (
	"context"
	"gin/app"
	"gin/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	dbCtx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	dbPool, err := db.ConnectDB(dbCtx, dbUrl)
	if err != nil {
		log.Fatal("Error while connecting to db")
	}
	defer dbPool.Close()

	server := app.CreateRouter(dbPool)
	server.Run(":8080")
}
