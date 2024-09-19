package middleware

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module("middleware", fx.Options(
		fx.Provide(NewJWTMiddleWare),
	))
}
