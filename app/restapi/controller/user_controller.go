package controller

import (
	"go-server-template/app/restapi/serializer"
	"go-server-template/internal/core/service"
	"go-server-template/internal/pkg/apires"
	"go-server-template/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UserController struct {
	logger      *logger.Logger
	UserService *service.UserService
}

type UserControllerParams struct {
	fx.In

	Logger      *logger.Logger
	UserService *service.UserService
}

func NewUserController(params UserControllerParams) *UserController {
	return &UserController{logger: params.Logger, UserService: params.UserService}
}

func (u *UserController) GetUsers(c *gin.Context) {
	u.logger.Info("UserController.GetUsers")

	users, err := u.UserService.GetUsers()
	if err != nil {
		apires.Error(c, http.StatusInternalServerError, "Failed to get users", err, "")
		return
	}

	apires.Success(c, "", users)
}

func (u *UserController) CreateUser(c *gin.Context) {
	u.logger.Info("UserController.CreateUser")

	var req serializer.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.logger.Errorf("Failed to bind request data: %v", err)
		apires.Error(c, http.StatusBadRequest, "Invalid request data", err, "")
		return
	}

	user, err := u.UserService.CreateUser(&req)
	if err != nil {
		u.logger.Errorf("Failed to create user: %v", err)
		apires.Error(c, http.StatusInternalServerError, "Failed to create user", err, "")
		return
	}

	apires.Success(c, "Created user successfully", user)
}
