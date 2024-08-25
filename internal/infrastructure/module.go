package infrastructure

import (
	"go-server-template/internal/infrastructure/http/router"

	"go.uber.org/fx"
)

// NewModule creates a new fx.Option with all the dependencies for the infrastructure layer
//
// Use this if you want to include all the infrastructure dependencies in your fx.App else use the individual.
func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			router.NewRouter,
		),
	)
}
