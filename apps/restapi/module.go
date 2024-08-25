package restapi

import (
	"go-server-template/apps/restapi/controllers"
	"go-server-template/apps/restapi/routes"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"restapi",
		controllers.NewModule(),
		routes.NewModule(),
	)
}
