package env

import "github.com/caarlos0/env/v9"

type AppEnv string

const (
	AppEnvDev AppEnv = "dev"
)

func Load(cfg interface{}) error {
	if err := env.Parse(cfg); err != nil {
		return err
	}
	return nil
}
