package controllers

import (
	"go-server-template/internal/pkg/apires"
	"log"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetUsers(c *gin.Context) {
	log.Println("UserController.GetUsers")

	apires.Success(c, "GetUsers", nil)
}
