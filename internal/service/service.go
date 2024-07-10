package service

import (
	"context"
	"totality_corp_kv/internal/domain"
	user "totality_corp_kv/internal/protogen/golang/users"
	"totality_corp_kv/internal/repository"

	"google.golang.org/grpc"
)

type User interface {
	CreateUser(ctx context.Context, id int64, user *domain.User) error
	AddUser(ctx context.Context, user *user.SingleUserRequest) (*user.Empty, error)
	GetUser(context.Context, *user.Id) (*user.SingleUserResponse, error)
	SearchUser(context.Context, *user.SearchFilter) (*user.GetUserListResponse, error)

	GetUserList(context.Context, *user.GetUserListRequest) (*user.GetUserListResponse, error)
}

type Services struct {
	User
}

type Deps struct {
	Repos *repository.Repositories
	*grpc.Server
}

func NewServices(deps Deps) *Services {
	userService := NewUserService(deps.Repos.User)
	user.RegisterUsersServer(deps.Server, userService)
	return &Services{
		User: userService,
	}
}
