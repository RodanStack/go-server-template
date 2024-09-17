package service

import (
	"go-server-template/db/sqlc"
	"go-server-template/internal/domain/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUsers() ([]sqlc.GetUsersRow, error) {
	users, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
