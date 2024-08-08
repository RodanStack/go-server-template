package logger

import (
	"errors"
	"go-server-template/pkg/config"
	"log"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	default:
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	err = zapLogger.Sync() // flushes buffer, if any
	if err != nil && !errors.Is(err, syscall.ENOTTY) {
		log.Fatalf("can't sync zap logger: %v", err)
	}

	return &Logger{zapLogger}
}
