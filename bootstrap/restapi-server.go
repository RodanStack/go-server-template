package bootstrap

import (
	"context"
	"errors"
	"go-server-template/apps/restapi"
	"go-server-template/internal/infrastructure"
	"go-server-template/internal/infrastructure/http/router"
	"go-server-template/pkg"
	"go-server-template/pkg/config"
	"go-server-template/pkg/logger"
	"log"
	"syscall"

	"go.uber.org/fx"
)

func RunRestAPIServer() {
	log.Println("Running the app")

	opt := fx.Options(
		pkg.NewModule(),
		infrastructure.NewModule(),
		restapi.NewModule(),
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

func startServer(lifecycle fx.Lifecycle, r *router.Router, env *config.Env, logger *logger.Logger) {
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
			// flushes buffer, if any
			if err := logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
				log.Fatalf("can't sync zap logger: %v", err)
			}
			log.Println("Stopping the server")
			return nil
		},
	})
}
