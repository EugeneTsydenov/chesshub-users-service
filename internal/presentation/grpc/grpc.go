package grpc

import (
	"context"
	proto "github.com/EugeneTsydenov/chesshub-users-service/internal/generated"
)

type UsersImplementation struct {
	proto.UnimplementedUsersServiceServer
}

func New() *UsersImplementation {
	return &UsersImplementation{}
}

func (i *UsersImplementation) Register(_ context.Context, req *proto.RegisterRequest) {

}
