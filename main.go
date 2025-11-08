package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	httpAdapter "github.com/not201/ninja-url/internal/adapters/http"
	"github.com/not201/ninja-url/internal/adapters/repositories"
	"github.com/not201/ninja-url/internal/core/services"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
)

//go:embed web/dist/*
var staticFS embed.FS

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system envs")
	}

	redis := redis.NewClient(&redis.Options{
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	repo := repositories.NewUrlRepository(redis)
	service := services.NewUrlService(repo)
	handler := httpAdapter.NewHandler(service, staticFS)

	r := httpAdapter.SetupRoutes(handler)

	http.ListenAndServe(os.Getenv("SERVER_PORT"), r)
}
