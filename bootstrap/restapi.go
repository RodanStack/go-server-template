package bootstrap

import (
	"context"
	"errors"
	"go-server-template/app/restapi"
	"go-server-template/db"
	"go-server-template/internal/core"
	"go-server-template/internal/infrastructure"
	"go-server-template/internal/infrastructure/http/router"
	"go-server-template/internal/infrastructure/persistence/database"
	"go-server-template/pkg"
	"go-server-template/pkg/config"
	"go-server-template/pkg/logger"
	"log"
	"syscall"

	"go.uber.org/fx"
)

func RunRestAPI() {
	log.Println("Running the app")

	opt := fx.Options(
		pkg.NewModule(),
		infrastructure.NewModule(),
		restapi.NewModule(),
		db.NewModule(),
		core.NewModule(),
		fx.Invoke(
			startServer,
		),
	)

	app := fx.New(opt, fx.NopLogger)

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

type StartServerParams struct {
	fx.In

	Router   *router.Router
	Env      *config.Env
	Logger   *logger.Logger
	Database *database.Database
}

func startServer(lifecycle fx.Lifecycle, params StartServerParams) {
	lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			params.Logger.Info("Starting the server")

			params.Database.RunMigrations()

			go func() {
				if err := params.Router.Run(":" + params.Env.ServerPort); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			params.Logger.Info("Stopping the server")
			params.Database.Close()
			// flushes buffer, if any
			if err := params.Logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
				log.Fatalf("can't sync zap logger: %v", err)
			}
			log.Println("Stopping the server")
			return nil
		},
	})
}
