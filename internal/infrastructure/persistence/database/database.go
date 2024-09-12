package database

import (
	"context"
	"fmt"
	"go-server-template/pkg/config"
	"log"
	"net"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(env *config.Env) *pgxpool.Pool {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s",
		env.DBUser, env.DBPassword, net.JoinHostPort(env.DBHost, env.DBPort)) + fmt.Sprintf("/%s?sslmode=disable", env.DBName)
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	return pool
}
