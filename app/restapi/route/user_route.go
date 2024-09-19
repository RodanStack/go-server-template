package route

import (
	"go-server-template/app/restapi/controller"
	"go-server-template/internal/infrastructure/http/middleware"
	"go-server-template/internal/infrastructure/http/router"

	"go.uber.org/fx"
)

type UserRoutes struct {
	router        *router.Router
	controller    *controller.UserController
	jwtMiddleware *middleware.JWTMiddleware
}

type userRoutesParams struct {
	fx.In

	Router        *router.Router
	Controller    *controller.UserController
	JWTMiddleware *middleware.JWTMiddleware
}

func NewUserRoutes(params userRoutesParams) *UserRoutes {
	return &UserRoutes{params.Router, params.Controller, params.JWTMiddleware}
}

func (u *UserRoutes) RegisterRoutes() {
	// Register routes here

	userRoutes := u.router.Group("users")
	userRoutes.POST("/login", u.controller.LoginUser)

	userRoutesWithAuth := u.router.Group("users", u.jwtMiddleware.Handle())
	userRoutesWithAuth.GET("", u.controller.GetUsers)
	userRoutesWithAuth.POST("", u.controller.CreateUser)
}
