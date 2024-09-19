package infrastructure

import (
	"go-server-template/internal/infrastructure/http/middleware"
	"go-server-template/internal/infrastructure/http/router"
	"go-server-template/internal/infrastructure/persistence/database"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			router.NewRouter,
			database.NewDatabase,
			middleware.NewModule,
		),
	)
}
