package repository

import (
	"go-server-template/db/sqlc"
	"go-server-template/internal/infrastructure/persistence/database"
	"go-server-template/pkg/logger"
)

type Repository struct {
	q      *sqlc.Queries
	db     *database.Database
	logger *logger.Logger
}

func NewRepository(db *database.Database, queries *sqlc.Queries, logger *logger.Logger) *Repository {
	return &Repository{
		q:      queries,
		db:     db,
		logger: logger,
	}
}
