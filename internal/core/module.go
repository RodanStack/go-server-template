package core

import (
	"go-server-template/internal/core/repository"
	"go-server-template/internal/core/service"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module("core", fx.Options(
		repository.NewModule(),
		service.NewModule(),
	))
}
