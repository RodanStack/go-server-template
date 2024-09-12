package database

import (
	"context"
	"fmt"
	"go-server-template/pkg/config"
	"go-server-template/pkg/logger"
	"net"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	conn   *pgxpool.Pool
	logger *logger.Logger
}

func NewDatabase(env *config.Env, logger *logger.Logger) *Database {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s",
		env.DBUser, env.DBPassword, net.JoinHostPort(env.DBHost, env.DBPort)) + fmt.Sprintf("/%s?sslmode=disable", env.DBName)
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		logger.Errorf("Unable to parse DATABASE_URL: %v\n", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Errorf("Unable to create connection pool: %v\n", err)
	}

	return &Database{
		conn:   pool,
		logger: logger,
	}
}

func (d *Database) Close() {
	if d.conn != nil {
		d.conn.Close()
		d.logger.Info("Database connection closed")
	}
}
