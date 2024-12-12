package app

import (
	"context"
	"fmt"
	"net"

	"github.com/EugeneTsydenov/chesshub-users-service/internal/config"
	usersv1 "github.com/EugeneTsydenov/chesshub-users-service/internal/generated/proto"
	"github.com/EugeneTsydenov/chesshub-users-service/internal/infrastructure/database"
	"github.com/EugeneTsydenov/chesshub-users-service/internal/presentation/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	cfg        *config.Config
	dbPool     database.Pool
	grpcServer *grpc.Server
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	dbPool, err := database.NewPool(ctx, cfg.Database.URL)

	if err != nil {
		return nil, fmt.Errorf("failed to initialize database pool: %w", err)
	}

	grpcServer := initGrpcServer()

	return &App{
		cfg:        cfg,
		dbPool:     dbPool,
		grpcServer: grpcServer,
	}, nil
}

func initGrpcServer() *grpc.Server {
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(server)
	usersService := v1.NewUsersService()
	usersv1.RegisterUsersServiceServer(server, usersService)
	return server
}

func (a *App) Run(_ context.Context) error {
	servicePort := "4040"
	listener, err := net.Listen("tcp", ":"+servicePort)

	if err != nil {
		return fmt.Errorf("failed to start listener on port %s: %w", servicePort, err)
	}

	fmt.Printf("gRPC server is running on port %s\n", servicePort)

	if err := a.grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("gRPC server terminated: %w", err)
	}

	return nil
}
