package service

import (
	"go-server-template/app/restapi/serializer"
	"go-server-template/db/sqlc"
	"go-server-template/internal/core/repository"
	"go-server-template/internal/pkg/auth"
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

func (s *UserService) CreateUser(userData *serializer.CreateUserRequest) (*sqlc.CreateUserRow, error) {
	s.logger.Info("UserService.CreateUser")

	hashedPassword, err := auth.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.CreateUser(&sqlc.CreateUserParams{
		Username: userData.Username,
		Password: hashedPassword,
		Email:    userData.Email,
		Status:   sqlc.UserStatusActive,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
