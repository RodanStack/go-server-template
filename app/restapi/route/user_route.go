package route

import (
	"go-server-template/app/restapi/controller"
	"go-server-template/internal/infrastructure/http/router"

	"go.uber.org/fx"
)

type UserRoutes struct {
	router     *router.Router
	controller *controller.UserController
}

type userRoutesParams struct {
	fx.In

	Router     *router.Router
	Controller *controller.UserController
}

func NewUserRoutes(params userRoutesParams) *UserRoutes {
	return &UserRoutes{params.Router, params.Controller}
}

func (u *UserRoutes) RegisterRoutes() {
	// Register routes here

	routes := u.router.Group("users")

	routes.GET("", u.controller.GetUsers)
	routes.POST("", u.controller.CreateUser)
}
