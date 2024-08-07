package app

import (
	"context"
	"go-server-template/internal/app/controllers"
	"go-server-template/internal/app/routes"
	"go-server-template/internal/infra"
	"go-server-template/internal/infra/router"
	"log"

	"go.uber.org/fx"
)

func Run() {
	log.Println("Running the app")

	opt := fx.Options(
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

func startServer(lifecycle fx.Lifecycle, r *router.Router) {
	lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := r.Engine.Run(":8080"); err != nil {
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
