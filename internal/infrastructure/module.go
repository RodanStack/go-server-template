package infrastructure

import (
	"go-server-template/internal/infrastructure/http"
	"go-server-template/internal/infrastructure/persistence"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module("infrastructure", fx.Options(
		http.NewModule(),
		persistence.NewModule(),
	))
}
