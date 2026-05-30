package user_postgres_repository

import (
	core_postgres_pool "github.com/sofifaro/to_go_list/internal/core/repository/postgres/pool"
)

type UsersRepository struct {
	pool core_postgres_pool.Pool
}

func NewUsersRepository(pool core_postgres_pool.Pool) *UsersRepository {
	return &UsersRepository{pool: pool}
}
