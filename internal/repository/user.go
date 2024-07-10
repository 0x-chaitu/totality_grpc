package repository

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"totality_corp_kv/internal/domain"
	user "totality_corp_kv/internal/protogen/golang/users"
	"totality_corp_kv/pkg/kvpair"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepo struct {
	kv *kvpair.Cache
}

func NewUserRepo(kv *kvpair.Cache) *UserRepo {
	return &UserRepo{
		kv: kv,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, id int64, input *domain.User) error {
	user, err := bson.Marshal(input)
	if err != nil {
		return err
	}
	r.kv.Set(id, user)
	return nil
}

func (r *UserRepo) AddUser(ctx context.Context, input *user.SingleUserRequest) error {
	user, err := bson.Marshal(domain.User{
		Id:      input.User.Id,
		Name:    input.User.Fname,
		City:    input.User.City,
		Phone:   input.User.Phone,
		Height:  input.User.Height,
		Married: input.User.Married,
	})
	if err != nil {
		return err
	}
	r.kv.Set(input.User.Id, user)
	return nil
}

func (r *UserRepo) GetUser(ctx context.Context, id int64) (*user.SingleUserResponse, error) {
	res, ok := r.kv.Get(id)
	if !ok {
		return nil, errors.New("no user found")
	}
	var u domain.User

	if err := bson.Unmarshal(res, &u); err != nil {
		return nil, err
	}
	return &user.SingleUserResponse{
		User: &user.User{
			Id:      u.Id,
			Fname:   u.Name,
			City:    u.City,
			Height:  u.Height,
			Phone:   u.Phone,
			Married: u.Married,
		},
	}, nil
}

func (r *UserRepo) GetUserList(context.Context) (*user.GetUserListResponse, error) {

	users := r.kv.GetAll()
	var res []*user.User
	var u domain.User
	for _, v := range users {
		if err := bson.Unmarshal(v, &u); err != nil {
			return nil, err
		}
		res = append(res, &user.User{
			Id:      u.Id,
			Fname:   u.Name,
			City:    u.City,
			Phone:   u.Phone,
			Height:  u.Height,
			Married: u.Married,
		})

	}

	return &user.GetUserListResponse{
		Users: res,
	}, nil
}

func (r *UserRepo) SearchUser(ctx context.Context, q string) (*user.GetUserListResponse, error) {
	users := r.kv.GetAll()
	var res []*user.User
	var u domain.User
	for _, v := range users {
		if err := bson.Unmarshal(v, &u); err != nil {
			return nil, err
		}
		if strings.Contains(u.Name, q) ||
			strings.Contains(u.City, q) ||
			strings.Contains(strconv.Itoa(int(u.Phone)), q) ||
			strings.Contains(strconv.FormatBool(u.Married), q) ||
			strings.Contains(strconv.Itoa(int(u.Height)), q) {
			res = append(res, &user.User{
				Id:      u.Id,
				Fname:   u.Name,
				City:    u.City,
				Phone:   u.Phone,
				Height:  u.Height,
				Married: u.Married,
			})
			continue
		}

	}

	return &user.GetUserListResponse{
		Users: res,
	}, nil
}
