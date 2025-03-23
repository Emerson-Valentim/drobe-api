package config

type Env struct {
	// Server configuration
	Port string `env:"PORT" envDefault:"8080"`

	// Postgres configuration
	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"postgres"`
	PostgresPort     string `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	PostgresDB       string `env:"POSTGRES_DB" envDefault:"drobe"`

	// S3/LocalStack configuration
	S3Endpoint string `env:"S3_ENDPOINT" envDefault:"http://localhost:4566"`
	S3Region   string `env:"S3_REGION" envDefault:"us-east-1"`
	S3Bucket   string `env:"S3_BUCKET" envDefault:"clothes"`
}
