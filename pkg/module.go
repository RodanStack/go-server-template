package pkg

import (
	"go-server-template/pkg/config"
	"go-server-template/pkg/logger"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(config.NewEnv),
		fx.Provide(logger.NewLogger),
	)
}
