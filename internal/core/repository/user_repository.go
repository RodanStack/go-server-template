package repository

import (
	"context"
	"go-server-template/db/sqlc"
)

type UserRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) *UserRepository {
	return &UserRepository{r}
}

func (r *UserRepository) GetUsers() ([]sqlc.GetUsersRow, error) {
	r.logger.Info("UserRepository.GetUsers")

	const pageSize = 10
	arg := sqlc.GetUsersParams{
		Limit:  pageSize,
		Offset: 0,
	}

	users, err := r.q.GetUsers(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) CreateUser(userParams *sqlc.CreateUserParams) (*sqlc.CreateUserRow, error) {
	r.logger.Info("UserRepository.CreateUser")

	user, err := r.q.CreateUser(context.Background(), *userParams)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
