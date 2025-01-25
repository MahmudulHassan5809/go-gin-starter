.PHONY: create-db drop-db migrate-up migrate-down help

# Run commands inside the app container
DOCKER_EXEC_APP = docker exec -it gin_app
DOCKER_EXEC_DB = docker exec -it postgres_db

create-db:
	@echo "Creating database: $(DB_NAME)..."
	@$(DOCKER_EXEC_DB) psql -U $(POSTGRES_USER) -d postgres -c "CREATE DATABASE $(DB_NAME) OWNER $(POSTGRES_USER);" 2>/dev/null || echo "Database already exists."
	@echo "Granting all privileges on database $(DB_NAME) to user $(DB_USER)..."
	@$(DOCKER_EXEC_DB) psql -U $(POSTGRES_USER) -d postgres -c "GRANT ALL PRIVILEGES ON DATABASE $(DB_NAME) TO $(DB_USER);"

drop-db:
	@echo "Dropping database: $(DB_NAME)..."
	@$(DOCKER_EXEC_DB) psql -U $(POSTGRES_USER) -d postgres -c "DROP DATABASE IF EXISTS $(DB_NAME);" || echo "Failed to drop the database."

migrate-up:
	@echo "Migrating up..."
	@$(DOCKER_EXEC_APP) migrate -path ./migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migrate-down:
	@echo "Migrating down..."
	@$(DOCKER_EXEC_APP) migrate -path ./migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

help:
	@echo "Available commands:"
	@echo "  make create-db    - Create the database with owner postgres and grant privileges to user."
	@echo "  make drop-db      - Drop the database."
	@echo "  make migrate-up   - Run migrations up using environment variables."
	@echo "  make migrate-down - Run migrations down using environment variables."
	@echo "  make help         - Show this help message."
