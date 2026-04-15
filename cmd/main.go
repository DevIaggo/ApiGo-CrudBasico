package main

import (
	"apiGo/internal/api"
	"apiGo/internal/config"
	"apiGo/internal/handler"
	"apiGo/internal/repository"
	"apiGo/internal/service"
	"context"
	"log"
	"net/http"
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

	repo := repository.NewRepository(db)
	svc := service.NewTaskService(repo)
	handler := handler.NewHandler(svc)

	r := api.SetupRoutes(handler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
