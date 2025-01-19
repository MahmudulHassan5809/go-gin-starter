# Load environment variables from .env file
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

.PHONY: create-db drop-db migrate-up migrate-down help


create-db:
	@echo "Creating database: $(DB_NAME)..."
	@PGPASSWORD=$(POSTGRES_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(POSTGRES_USER) -d postgres -c "CREATE DATABASE $(DB_NAME) OWNER postgres;" 2>/dev/null || echo "Database already exists."
	@echo "Granting all privileges on database $(DB_NAME) to user $(DB_USER)..."
	@PGPASSWORD=$(POSTGRES_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(POSTGRES_USER) -d postgres -c "GRANT ALL PRIVILEGES ON DATABASE $(DB_NAME) TO $(DB_USER);"


drop-db:
	@echo "Dropping database: $(DB_NAME)..."
	@PGPASSWORD=$(POSTGRES_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(POSTGRES_USER) -d postgres -c "DROP DATABASE IF EXISTS $(DB_NAME);" || echo "Failed to drop the database."


migrate-up:
	@echo "Migrating up..."
	@migrate -path ./migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up


migrate-down:
	@echo "Migrating down..."
	@migrate -path ./migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down


help:
	@echo "Available commands:"
	@echo "  make create-db    - Create the database with owner postgres and grant privileges to user."
	@echo "  make drop-db      - Drop the database."
	@echo "  make migrate-up   - Run migrations up using environment variables."
	@echo "  make migrate-down - Run migrations down using environment variables."
	@echo "  make help         - Show this help message."