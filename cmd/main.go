package main

import (
	"apiGo/internal/config"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ctx := context.Background()
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("Credentials of databas is invalid")
	}
	db, err := config.NewConnectionPool(ctx, connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connected to database")
}
