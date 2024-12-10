package app

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-users-service/internal/config"
	"log"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	return a.initDependencies(ctx)
}

func (a *App) initDependencies(_ context.Context) error {
	c, err := config.Load()
	log.Print(c)
	return err
}
