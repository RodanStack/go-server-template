package logger

import (
	"errors"
	"fmt"
	"go-server-template/pkg/config"
	"log"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sys/unix"
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

	zapLogger, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	return &Logger{zapLogger}
}

func (l *Logger) Sync() error {
	// Try syncing and handle errors like ENOTTY gracefully
	err := l.Logger.Sync()
	if err != nil && !errors.Is(err, unix.ENOTTY) && !errors.Is(err, syscall.ENOTTY) {
		return err
	}

	return nil
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}
