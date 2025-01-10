include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

# Run
.PHONY: run-gateway
run-gateway:
	@go run ./gateway/cmd

.PHONY: run-randomizer
run-randomizer:
	@go run ./randomizer/cmd

.PHONY: run
run:
	@go run ./gateway/cmd & go run ./randomizer/cmd

# Tests
.PHONY: test
test:
	@go test -v ./...

# Migrations
.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_URI) up

.PHONY: migrate-down
migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_URI) down $(filter-out $@,$(MAKECMDGOALS))

# Docs
.PHONY: gen-docs
gen-docs:
	@swag init -g ./gateway/cmd/main.go -o ./gateway/docs && swag fmt

# Help
.PHONY: help
help:
	@echo "Usage: make <command>"
	@echo ""
	@echo "Commands:"
	@echo "  run-gateway     Run gateway service"
	@echo "  run-randomizer  Run randomizer service"
	@echo "  run             Run both services"
	@echo "  test            Run tests"
	@echo "  migrate-create  Create new migration"
	@echo "  migrate-up      Apply all migrations"
	@echo "  migrate-down    Rollback migrations"
	@echo "  gen-docs        Generate swagger docs"
	@echo "  help            Show this help message"
	@echo ""
	@echo "For more information, see the README.md file."
	@echo ""
	@echo "Examples:"
	@echo "  make run"
	@echo "  make test"
	@echo "  make migrate-create name=create_users_table"
	@echo "  make migrate-up"
	@echo "  make gen-docs"
	@echo "  make help"
