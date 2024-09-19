package database

import (
	"context"
	"fmt"
	"go-server-template/pkg/config"
	"go-server-template/pkg/logger"
	"net"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	goose "github.com/pressly/goose/v3"

	"github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	*pgxpool.Pool
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
		Pool:   pool,
		logger: logger,
	}
}

func (d *Database) Close() {
	if d.Pool != nil {
		d.Pool.Close()
		d.logger.Info("Database connection closed")
	}
}

func (d *Database) RunMigrations() {
	d.logger.Info("Running migrations")

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	db := stdlib.OpenDBFromPool(d.Pool)
	if err := goose.Up(db, "db/migrations"); err != nil {
		panic(err)
	}
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// alternative way to run migrations
// func (d *Database) RunMigrations() {
// 	d.logger.Info("Running migrations")
// 	db, err := goose.OpenDBWithDriver("postgres", d.dbURL)
// 	if err != nil {
// 		d.logger.Errorf("Unable to open database: %v\n", err)
// 	}
// 	defer db.Close()

// 	if err = goose.Up(db, "db/migrations"); err != nil {
// 		d.logger.Errorf("Unable to apply migrations: %v\n", err)
// 	}

// 	d.logger.Info("Migrations applied")
// }
