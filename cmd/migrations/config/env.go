package config

type Env struct {
	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"postgres"`
	PostgresPort     string `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	PostgresDB       string `env:"POSTGRES_DB" envDefault:"drobe"`
}
