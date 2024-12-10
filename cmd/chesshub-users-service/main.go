package main

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-users-service/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	a := app.New()
	err := a.Run(ctx)

	if err != nil {
		log.Fatalf("Application running with error: %v", err)
	}

	log.Print("Application running successfully")
}
