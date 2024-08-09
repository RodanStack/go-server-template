package infrastructure

import (
	"go-server-template/internal/infrastructure/router"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			router.NewRouter,
		),
	)
}
