package logger

import (
	"go-server-template/pkg/config"
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(env *config.Env) *Logger {
	var zapConfig zap.Config

	environment := env.Environment

	switch environment {
	case "production":
		zapConfig = zap.NewProductionConfig()
	case "development":
		zapConfig = zap.NewDevelopmentConfig()
	default:
		zapConfig = zap.NewDevelopmentConfig()
	}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	defer func() {
		err = zapLogger.Sync() // flushes buffer, if any
		if err != nil {
			log.Fatalf("can't sync zap logger: %v", err)
		}
	}()

	return &Logger{zapLogger}
}
