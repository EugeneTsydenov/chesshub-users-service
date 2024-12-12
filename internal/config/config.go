package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	URL string `env:"DATABASE_URL, required"`
}

func Load(ctx context.Context) (*Config, error) {
	var c Config

	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	return &c, nil
}
