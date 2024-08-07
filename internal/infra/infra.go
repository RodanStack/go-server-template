package infra

import (
	"go-server-template/internal/infra/router"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			router.NewRouter,
		),
	)
}
