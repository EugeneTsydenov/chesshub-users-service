package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

type DatabaseConfig struct {
	DatabaseUrl string `env:"DATABASE_URL, required"`
}

type Config struct {
	Database *DatabaseConfig
}

func Load() (*Config, error) {
	ctx := context.Background()

	var c *Config

	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, err
	}

	return c, nil
}
