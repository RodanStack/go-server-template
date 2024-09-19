package restapi

import (
	"go-server-template/app/restapi/controller"
	"go-server-template/app/restapi/route"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"restapi",
		controller.NewModule(),
		route.NewModule(),
	)
}
