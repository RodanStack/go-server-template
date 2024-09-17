# Include environment variables from .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

# Tool
GOOSE_BIN=docker compose exec web goose
SQLC_BIN=docker compose exec web sqlc

#------------------------------------------------

# For Migrations

#------------------------------------------------

# Database connection details
DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Migration directory
MIGRATIONS_DIR = db/migrations

# Show migration status
.PHONY: migrate-status
migrate-status:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status

# Create a new migration with a user-provided name
.PHONY: create-migration
create-migration:
	@read -p "What is the name of migration? " NAME; \
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) create $$NAME sql

# Apply migrations
.PHONY: migrate-up
migrate-up:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

# Rollback the last migration
.PHONY: migrate-down
migrate-down:
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

#-------------------------------------------------

# For SQLC

#-------------------------------------------------

SQLC_CONFIG_LOCATION = sqlc.yaml

# Generate SQLC
.PHONY: sqlc-generate
sqlc-generate:
	$(SQLC_BIN) generate -f $(SQLC_CONFIG_LOCATION)

# Help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make migrate-status   - Show migration status"
	@echo "  make create-migration - Create a new migration (e.g., make create-migration)"
	@echo "  make migrate-up       - Apply all migrations"
	@echo "  make migrate-down     - Rollback the last migration"
	@echo "  make help             - Display this help message"
	@echo "  make sqlc-generate   - Generate SQLC"
