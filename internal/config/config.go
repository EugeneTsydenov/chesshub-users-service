package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	DatabaseUrl string `env:"DATABASE_URL, required"`
}

func Load() (*Config, error) {
	ctx := context.Background()

	var c Config

	if err := envconfig.Process(ctx, &c); err != nil {
		return &c, err
	}

	return &c, nil
}
