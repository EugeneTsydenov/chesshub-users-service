package main

import (
	"context"
	"log"

	"github.com/EugeneTsydenov/chesshub-users-service/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatalf("Application terminated with error: %v", err)
	}

	log.Print("Application terminated successfully")
}
