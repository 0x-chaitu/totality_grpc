package repository

import (
	"context"
	"totality_corp_kv/internal/domain"
	user "totality_corp_kv/internal/protogen/golang/users"
	"totality_corp_kv/pkg/kvpair"
)

type User interface {
	CreateUser(ctx context.Context, id int64, user *domain.User) error
	AddUser(ctx context.Context, user *user.SingleUserRequest) error
	GetUser(ctx context.Context, id int64) (*user.SingleUserResponse, error)
	GetUserList(context.Context) (*user.GetUserListResponse, error)
	SearchUser(ctx context.Context, inp string) (*user.GetUserListResponse, error)
}

type Repositories struct {
	User
}

func NewRepositories(kv *kvpair.Cache) *Repositories {
	return &Repositories{
		User: NewUserRepo(kv),
	}
}
