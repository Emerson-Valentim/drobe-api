package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/emersonvalentim/drobe-api/cmd/migrations/config"
	"github.com/emersonvalentim/drobe-api/internal/env"
)

func main() {
	var cfg config.Env
	if err := env.Load(&cfg); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		os.Exit(1)
	}

	migrator, err := migrate.New("file://db/migrations", getDatabaseURL(cfg))
	if err != nil {
		log.Fatalf("Failed to create migrator: %v", err)
		os.Exit(1)
	}

	if err := migrator.Up(); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
		os.Exit(1)
	}
}

func getDatabaseURL(cfg config.Env) string {
	return "postgres://" + cfg.PostgresUser + ":" + cfg.PostgresPassword + "@" + cfg.PostgresHost + ":" + cfg.PostgresPort + "/" + cfg.PostgresDB + "?sslmode=disable"
}
