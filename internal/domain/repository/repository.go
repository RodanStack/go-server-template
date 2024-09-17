package repository

import (
	"go-server-template/db/sqlc"
	"go-server-template/internal/infrastructure/persistence/database"
)

type Repository struct {
	q  *sqlc.Queries
	db *database.Database
}

func NewRepository(db *database.Database, queries *sqlc.Queries) *Repository {

	return &Repository{
		q:  queries,
		db: db,
	}
}
