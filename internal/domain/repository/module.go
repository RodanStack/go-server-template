package repository

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(NewRepository),
		fx.Provide(NewUserRepository),
	)
}
