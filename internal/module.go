package internal

import (
	"go-server-template/internal/core"
	"go-server-template/internal/infrastructure"
	"go-server-template/internal/pkg"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module("internal",
		fx.Options(
			pkg.NewModule(),
			core.NewModule(),
			infrastructure.NewModule(),
		),
	)
}
