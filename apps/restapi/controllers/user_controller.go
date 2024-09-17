package controllers

import (
	"go-server-template/internal/domain/service"
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
