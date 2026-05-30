package users_service

import (
	"context"

	"github.com/sofifaro/to_go_list/internal/core/domain"
)

type UsersService struct {
	usersReposiroty UsersReposiroty
}

type UsersReposiroty interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

func NewUsersService(usersReposiroty UsersReposiroty) *UsersService {
	return &UsersService{
		usersReposiroty: usersReposiroty,
	}
}
