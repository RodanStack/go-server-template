package app

import (
	"context"
	"go-server-template/internal/app/controllers"
	"go-server-template/internal/app/routes"
	"go-server-template/internal/infra"
	"go-server-template/internal/infra/router"
	"go-server-template/pkg/config"
	"go-server-template/pkg/logger"
	"log"

	"go.uber.org/fx"
)

func Run() {
	log.Println("Running the app")

	opt := fx.Options(
		fx.Provide(config.NewEnv),
		fx.Provide(logger.NewLogger),
		controllers.NewModule(),
		routes.NewModule(),
		infra.NewModule(),
		fx.Invoke(
			startServer,
		),
	)

	app := fx.New(opt)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := app.Stop(context.Background()); err != nil {
			log.Fatal(err)
		}

		log.Println("App stopped")
	}()

	<-app.Done()

	log.Println("App done")
}

func startServer(lifecycle fx.Lifecycle, r *router.Router, env *config.Env) {
	lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := r.Run(":" + env.ServerPort); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Println("Stopping the server")
			return nil
		},
	})
}
