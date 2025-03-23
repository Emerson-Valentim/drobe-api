package main

import (
	"context"
	"log"

	"github.com/emersonvalentim/drobe-api/cmd/api/config"
	"github.com/emersonvalentim/drobe-api/cmd/api/router"
	"github.com/emersonvalentim/drobe-api/internal/env"
	"github.com/emersonvalentim/drobe-api/internal/postgres"
	"github.com/emersonvalentim/drobe-api/services/auth"
	"github.com/emersonvalentim/drobe-api/services/inventory"
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

	inventoryService := inventory.NewService(inventory.NewRepository(postgres))
	authService := auth.NewService(cfg.AuthSecret, cfg.JWTSecret, auth.NewRepository(postgres))

	router := router.NewRouter(inventoryService, authService, &cfg)
	router.Register()
}

func getDatabaseURL(cfg config.Env) string {
	return "postgres://" + cfg.PostgresUser + ":" + cfg.PostgresPassword + "@" + cfg.PostgresHost + ":" + cfg.PostgresPort + "/" + cfg.PostgresDB + "?sslmode=disable"
}
