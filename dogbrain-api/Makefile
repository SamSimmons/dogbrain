# Simple Makefile for a Go project

# Include .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

# Database migrations
check-db-url:
ifndef DATABASE_URL
	$(error DATABASE_URL is not set. Please set it to your database connection string)
endif

migrate-up: check-db-url
	@echo "Running migrations..."
	migrate -path internal/db/migrations -database "${DATABASE_URL}" up

migrate-down: check-db-url
	@echo "Rolling back migrations..."
	migrate -path internal/db/migrations -database "${DATABASE_URL}" down

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir internal/db/migrations -seq $$name

migrate-force: check-db-url
	@read -p "Enter version to force: " version; \
	migrate -path internal/db/migrations -database "${DATABASE_URL}" force $$version

migrate-version: check-db-url
	@echo "Current migration version:"
	@migrate -path internal/db/migrations -database "${DATABASE_URL}" version