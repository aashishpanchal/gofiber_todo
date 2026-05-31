# https://just.systems

# Load .env
set dotenv-load

# Cross-platform shell (Windows: Git Bash)
set windows-shell := ["bash", "-cu"]

# Project Info
PROJECT_NAME := "todo_list"
MAIN_PATH := "./cmd/main.go"
BINARY_NAME := "./cmd/main.exe"

# Goose Config
GOOSE_DRIVER := "pgx"
GOOSE_DIR := "./sqlpg/schema"
GOOSE_DBSTRING := env("DB_URL")

# Default recipe - show help
default:
    @just --list

# Start development server with live reload (Air)
dev:
    @if command -v air >/dev/null 2>&1; then \
        air; \
    else \
        echo "Air not found. Install with:"; \
        echo "go install github.com/air-verse/air@latest"; \
        exit 1; \
    fi

# Run the production binary (requires build first)
start:
    @./{{BINARY_NAME}}

# Build optimized binary.
build:
    @go build -ldflags="-w -s" -trimpath -o {{BINARY_NAME}} {{MAIN_PATH}}
    @echo "Build complete: {{BINARY_NAME}}"
    @ls -lh {{BINARY_NAME}} | awk '{print "Binary size: " $5}'

# Run all tests with verbose output
code-test:
    @go test ./tests/... -v

# Format all Go source code
code-format:
    @go fmt ./...

# Download and cache all dependencies
mod-deps:
    @go mod download

# Remove unused dependencies and clean go.mod
mod-tidy:
    @go mod tidy

# Remove build artifacts and temporary files
mod-clean:
    @rm -rf bin/ tmp/ coverage.out coverage.html

# Run all pending migrations [GOOSE_DBSTRING=<db_url>]
migrate-up:
    @if command -v goose >/dev/null 2>&1; then \
        GOOSE_DRIVER={{GOOSE_DRIVER}} GOOSE_DBSTRING="{{GOOSE_DBSTRING}}" goose -dir {{GOOSE_DIR}} up; \
    else \
        echo "goose not found. Install with:"; \
        echo "go install github.com/pressly/goose/v3/cmd/goose@latest"; \
        exit 1; \
    fi

# Rollback last migration [GOOSE_DBSTRING=<db_url>]
migrate-down:
    @GOOSE_DRIVER={{GOOSE_DRIVER}} GOOSE_DBSTRING="{{GOOSE_DBSTRING}}" goose -dir {{GOOSE_DIR}} down

# Show migration status [GOOSE_DBSTRING=<db_url>]
migrate-status:
    @GOOSE_DRIVER={{GOOSE_DRIVER}} GOOSE_DBSTRING="{{GOOSE_DBSTRING}}" goose -dir {{GOOSE_DIR}} status

# Rollback all migrations [GOOSE_DBSTRING=<db_url>]
migrate-reset:
    @GOOSE_DRIVER={{GOOSE_DRIVER}} GOOSE_DBSTRING="{{GOOSE_DBSTRING}}" goose -dir {{GOOSE_DIR}} reset

# Generate type-safe Go code from SQL queries
sqlc:
    @if command -v sqlc >/dev/null 2>&1; then \
        cd sqlpg && sqlc generate; \
    else \
        echo "sqlc not found. Install with:"; \
        echo "go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest"; \
        exit 1; \
    fi
