package restapi

import (
	"go-server-template/app/restapi/controllers"
	"go-server-template/app/restapi/routes"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"restapi",
		controllers.NewModule(),
		routes.NewModule(),
	)
}
