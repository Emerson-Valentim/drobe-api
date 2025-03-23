package main

import (
	"context"
	"log"

	"github.com/emersonvalentim/drobe-api/cmd/api/config"
	"github.com/emersonvalentim/drobe-api/cmd/api/router"
	"github.com/emersonvalentim/drobe-api/internal/env"
	"github.com/emersonvalentim/drobe-api/internal/postgres"
	"github.com/emersonvalentim/drobe-api/services/clothing"
)

func main() {
	var cfg config.Env
	// Load configuration
	if err := env.Load(&cfg); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	postgres, err := postgres.New(context.Background(), getDatabaseURL(cfg))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	clothService := clothing.NewService(clothing.NewRepository(postgres))

	router := router.NewRouter(clothService, &cfg)
	router.Register()
}

func getDatabaseURL(cfg config.Env) string {
	return "postgres://" + cfg.PostgresUser + ":" + cfg.PostgresPassword + "@" + cfg.PostgresHost + ":" + cfg.PostgresPort + "/" + cfg.PostgresDB + "?sslmode=disable"
}
