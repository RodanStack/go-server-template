package route

import "go.uber.org/fx"

type Collection struct {
	fx.In

	UserRoutes *UserRoutes
}

type Routes interface {
	RegisterRoutes()
}

func NewRoutes(c Collection) []Routes {
	return []Routes{
		c.UserRoutes,
	}
}

func RegisterRoutes(routes []Routes) {
	for _, r := range routes {
		r.RegisterRoutes()
	}
}

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			NewUserRoutes,
			NewRoutes,
		),
		fx.Invoke(RegisterRoutes),
	)
}
