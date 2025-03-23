package main

import (
	"log"

	"github.com/emersonvalentim/drobe-api/cmd/api/config"
	"github.com/emersonvalentim/drobe-api/cmd/api/router"
	"github.com/emersonvalentim/drobe-api/internal/env"
	"github.com/emersonvalentim/drobe-api/services/clothing"
)

func main() {
	var cfg config.Env
	// Load configuration
	if err := env.Load(&cfg); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	clothService := clothing.NewService()

	router := router.NewRouter(clothService, &cfg)
	router.Register()
}
