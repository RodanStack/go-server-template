package bootstrap

import (
	"context"

	"go-server-template/internal/infrastructure"
	"go-server-template/internal/infrastructure/http/router"
	"go-server-template/internal/restapi/controllers"
	"go-server-template/internal/restapi/routes"
	"go-server-template/pkg"
	"go-server-template/pkg/config"
	"log"

	"go.uber.org/fx"
)

func RunRestAPIServer() {
	log.Println("Running the app")

	opt := fx.Options(
		pkg.NewModule(),
		infrastructure.NewModule(),
		controllers.NewModule(),
		routes.NewModule(),
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