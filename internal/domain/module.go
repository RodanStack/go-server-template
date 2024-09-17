package domain

import (
	"go-server-template/internal/domain/repository"
	"go-server-template/internal/domain/service"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module("domain", fx.Options(
		repository.NewModule(),
		service.NewModule(),
	))
}
