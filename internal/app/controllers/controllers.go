package controllers

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			NewUserController,
		),
	)
}
