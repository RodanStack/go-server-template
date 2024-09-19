package pkg

import (
	"go-server-template/internal/pkg/auth"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module("pkg",
		fx.Options(
			auth.NewModule(),
		))
}
