package users_service

import (
	"context"
	"fmt"

	"github.com/sofifaro/to_go_list/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user: %w", err)
	}

	user, err := s.usersReposiroty.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user in repository: %w", err)
	}

	return user, nil
}
