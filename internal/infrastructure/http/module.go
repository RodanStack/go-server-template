package http

import (
	"go-server-template/internal/infrastructure/http/middleware"
	"go-server-template/internal/infrastructure/http/router"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			router.NewRouter,
		),
		middleware.NewModule(),
	)
}
