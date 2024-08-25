package controllers

import (
	"go-server-template/internal/pkg/apires"
	"go-server-template/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UserController struct {
	logger *logger.Logger
}

type UserControllerParams struct {
	fx.In

	Logger *logger.Logger
}

func NewUserController(params UserControllerParams) *UserController {
	return &UserController{logger: params.Logger}
}

func (u *UserController) GetUsers(c *gin.Context) {
	u.logger.Info("UserController.GetUsers")

	apires.Success(c, "GetUsers", nil)
}
