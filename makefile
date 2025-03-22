# Load environment variables from the .env file
include .env
export $(shell sed 's/=.*//' .env)


# Database URL to use in migration command
DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/todo?sslmode=disable



# Migration path
MIGRATION_PATH=./migrations

# Command to run migrations
migrate:
	@echo "Running migrations..."
	@migrate -path $(MIGRATION_PATH) -database $(DB_URL) up

# Rollback the migrations
rollback:
	@echo "Rolling back migrations..."
	@migrate -path $(MIGRATION_PATH) -database $(DB_URL) down 1

# Reset the database (rollback all migrations)
reset:
	@echo "Rolling back all migrations..."
	@migrate -path $(MIGRATION_PATH) -database $(DB_URL) drop



.PHONY: all build run clean

all: build

build:
	go build -o bin/main cmd/main.go

run: build
	./bin/main

clean:
	rm -f bin/main