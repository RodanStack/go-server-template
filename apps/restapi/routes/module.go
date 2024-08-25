package routes

import "go.uber.org/fx"

type RouteCollection struct {
	fx.In

	UserRoutes *UserRoutes
}

type Routes interface {
	RegisterRoutes()
}

func NewRoutes(rc RouteCollection) []Routes {
	return []Routes{
		rc.UserRoutes,
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
