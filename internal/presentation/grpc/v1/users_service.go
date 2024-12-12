package v1

import (
	"context"

	usersv1 "github.com/EugeneTsydenov/chesshub-users-service/internal/generated/proto"
)

type UsersService struct {
	usersv1.UnimplementedUsersServiceServer
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) Register(ctx context.Context, req *usersv1.RegisterRequest) (*usersv1.RegisterResponse, error) {
	// TODO: Delegate to use case layer
	return &usersv1.RegisterResponse{
		Message: "Hello, " + req.Username,
	}, nil
}
