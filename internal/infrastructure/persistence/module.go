package persistence

import (
	"go-server-template/internal/infrastructure/persistence/database"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			database.NewDatabase,
		),
	)
}
