package db

import (
	"go-server-template/db/sqlc"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"db",
		sqlc.NewModule(),
	)
}
