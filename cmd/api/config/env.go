package config

import (
	"github.com/emersonvalentim/drobe-api/internal/env"
)

type Env struct {
	AppEnv env.AppEnv `env:"APP_ENV" envDefault:"dev"`
	// Server configuration
	Port string `env:"PORT" envDefault:"8080"`

	AWS struct {
		AccessKeyID     string `env:"AWS_ACCESS_KEY_ID" envDefault:""`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY" envDefault:""`
		Region          string `env:"AWS_REGION" envDefault:"us-east-1"`
	}

	// Postgres configuration
	Postgres struct {
		Host     string `env:"POSTGRES_HOST" envDefault:"postgres"`
		Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
		User     string `env:"POSTGRES_USER" envDefault:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
		DB       string `env:"POSTGRES_DB" envDefault:"drobe"`
	}

	// S3 configuration
	S3 struct {
		Endpoint string `env:"S3_ENDPOINT" envDefault:"http://localstack:4566"`
		Bucket   string `env:"S3_BUCKET" envDefault:"inventory"`
	}

	// JWT configuration
	JWTSecret string `env:"JWT_SECRET" envDefault:"secret"`

	// Auth configuration
	AuthSecret string `env:"AUTH_SECRET" envDefault:"secret"`
}
