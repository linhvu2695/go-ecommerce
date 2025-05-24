APP_NAME = server

GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING = "root:root1234@tcp(127.0.0.1:3306)/db"
GOOSE_MIGRATION_DIR ?= sql/schema

dev:
	go run ./cmd/${APP_NAME}

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: Migration name is required. Usage: make migrate-create name=your_migration_name"; \
		exit 1; \
	fi
	@goose -dir=$(GOOSE_MIGRATION_DIR) create "$(name)" sql
migrate-up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
migrate-down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
migrate-reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlcgen:
	sqlc generate
wire:
	cd internal/wire && wire
swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

docker-down:
	docker compose down
docker-up:
	docker compose -f environment/docker-compose-dev.yaml up -d
docker-buildup:
	docker compose up --build -d