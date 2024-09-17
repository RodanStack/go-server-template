package service

import (
	"go-server-template/db/sqlc"
	"go-server-template/internal/core/repository"
	"go-server-template/pkg/logger"
)

type UserService struct {
	logger   *logger.Logger
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository, logger *logger.Logger) *UserService {
	return &UserService{userRepo: userRepo, logger: logger}
}

func (s *UserService) GetUsers() ([]sqlc.GetUsersRow, error) {
	s.logger.Info("UserService.GetUsers")

	users, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
