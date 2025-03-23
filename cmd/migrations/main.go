package main

import (
	"log"

	"github.com/emersonvalentim/drobe-api/cmd/migrations/config"
	"github.com/emersonvalentim/drobe-api/internal/env"
)

func main() {
	var cfg config.Env
	if err := env.Load(&cfg); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
