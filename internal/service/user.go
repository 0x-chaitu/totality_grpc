package service

import (
	"context"
	"totality_corp_kv/internal/domain"
	user "totality_corp_kv/internal/protogen/golang/users"
	"totality_corp_kv/internal/repository"
)

type UserService struct {
	repo repository.User
	user.UnimplementedUsersServer
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, id int64, user *domain.User) error {
	return s.repo.CreateUser(ctx, id, user)
}

func (s *UserService) AddUser(ctx context.Context, user *user.SingleUserRequest) (*user.Empty, error) {
	return nil, s.repo.AddUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, inp *user.Id) (*user.SingleUserResponse, error) {
	return s.repo.GetUser(ctx, inp.Id)
}

func (s *UserService) SearchUser(ctx context.Context, inp *user.SearchFilter) (*user.GetUserListResponse, error) {
	return s.repo.SearchUser(ctx, inp.Query)
}

func (s *UserService) GetUserList(ctx context.Context, user *user.GetUserListRequest) (*user.GetUserListResponse, error) {
	return s.repo.GetUserList(ctx)
}
