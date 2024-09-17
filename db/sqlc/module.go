package sqlc

import (
	"go-server-template/internal/infrastructure/persistence/database"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(newQueries),
	)
}

func newQueries(db *database.Database) *Queries {
	// checking if db implements DBTX interface
	var _ DBTX = db

	return New(db)
}
