package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/emersonvalentim/drobe-api/cmd/api/config"
	"github.com/emersonvalentim/drobe-api/cmd/api/router"
	"github.com/emersonvalentim/drobe-api/internal/env"
	"github.com/emersonvalentim/drobe-api/internal/filestore"
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

	awsConfig, err := setupAWSConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to setup AWS config: %v", err)
	}

	postgres, err := postgres.New(context.Background(), getDatabaseURL(cfg))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	imageBucketClient := filestore.NewS3Client(awsConfig, cfg.S3.Bucket, cfg.S3.Endpoint)

	inventoryService := inventory.NewService(inventory.NewRepository(postgres), imageBucketClient)
	authService := auth.NewService(cfg.AuthSecret, cfg.JWTSecret, auth.NewRepository(postgres))

	router := router.NewRouter(inventoryService, authService, &cfg)
	router.Register()
}

func getDatabaseURL(cfg config.Env) string {
	return "postgres://" + cfg.Postgres.User + ":" + cfg.Postgres.Password + "@" + cfg.Postgres.Host + ":" + cfg.Postgres.Port + "/" + cfg.Postgres.DB + "?sslmode=disable"
}

func setupAWSConfig(cfg config.Env) (aws.Config, error) {
	awsConfig, err := awsconfig.LoadDefaultConfig(context.Background(), awsconfig.WithRegion(cfg.AWS.Region))
	if err != nil {
		return aws.Config{}, err
	}

	if cfg.AppEnv == env.AppEnvDev {
		awsConfig.Credentials = credentials.NewStaticCredentialsProvider(cfg.AWS.AccessKeyID, cfg.AWS.SecretAccessKey, "")
	}

	return awsConfig, nil
}
