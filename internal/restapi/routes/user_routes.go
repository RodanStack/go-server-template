package routes

import (
	"go-server-template/internal/infrastructure/http/router"
	"go-server-template/internal/restapi/controllers"

	"go.uber.org/fx"
)

type UserRoutes struct {
	router     *router.Router
	controller *controllers.UserController
}

type userRoutesParams struct {
	fx.In

	Router     *router.Router
	Controller *controllers.UserController
}

func NewUserRoutes(params userRoutesParams) *UserRoutes {
	return &UserRoutes{params.Router, params.Controller}
}

func (u *UserRoutes) RegisterRoutes() {
	// Register routes here

	routes := u.router.Group("/api/v1/users")

	routes.GET("/", u.controller.GetUsers)
}
