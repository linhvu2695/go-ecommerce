# Boilerplate for Golang Applications

- Run in dev: `make dev`
- Run in containers: `make docker-buildup`

## Setup
For first time setup:
- Start the containers: `make docker-buildup`
- Run migrations: `make migrate-up`

## Configurations
- Defined in `config/`
- `github.com/spf13/viper`: configuration Go library
- Represent in code as `pkg/settings/section.go`

## Dependency Injection
- `github.com/google/wire`: dependency injection Go library
- Define dependencies in `wire/`
- Command `wire`: init the controllers with all its dependency injected

## Log
- `go.uber.org/zap`: super fast logging library
- Stored in `storage/logs/`

## Database
- `goose`: database versioning tool
    - Versioning sql is in `sql/schema/`
    - Run `goose up` or `goose down` to apply or revert migrations
- `sqlc`: SQL to Go source code generator
    - Generates Go structs from SQL migrations in `sql/queries/`
    - Generate by command `sqlc generate`
    - This package effectively remove the need for repository layer

## Email
- `net/smtp`: Go's built-in SMTP client
- `templates/email/`: HTML templates for emails. Use this to send email with proper layout and style.

## Documentation
- Generate swagger documentation: `make swag`
- Access at: http://localhost:8082/swagger/index.html