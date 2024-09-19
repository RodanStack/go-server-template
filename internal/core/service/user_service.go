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
	jwtAuth  *auth.JWT
}

func NewUserService(userRepo *repository.UserRepository, logger *logger.Logger, jwtAuth *auth.JWT) *UserService {
	return &UserService{userRepo: userRepo, logger: logger, jwtAuth: jwtAuth}
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

func (s *UserService) LoginUser(userData *serializer.LoginUserRequest) (string, error) {
	s.logger.Info("UserService.LoginUser")

	user, err := s.userRepo.GetRawUserByUsername(userData.Username)
	if err != nil {
		return "", err
	}

	if !auth.CheckPasswordHash(userData.Password, user.Password) {
		return "", auth.ErrInvalidCredentials
	}

	generatedToken, err := s.jwtAuth.Generate(user.ID)
	if err != nil {
		return "", err
	}

	return generatedToken, nil
}
