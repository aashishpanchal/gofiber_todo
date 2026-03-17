# TodoList Makefile
# Email: aipanchal51@gmail.com
# Author: Aashish Panchal

.DEFAULT_GOAL := help

# Output Colors
RED    := \033[0;31m
GREEN  := \033[0;32m
YELLOW := \033[1;33m
BLUE   := \033[0;34m
CYAN   := \033[0;36m
NC     := \033[0m

# Project Info
PROJECT_NAME := todo_list
MAIN_PATH := ./cmd/main.go
BINARY_NAME :=./cmd/main.bin
# Goose Config
GOOSE_DRIVER  := pgx
GOOSE_DIR     := ./sqlpg/schema
GOOSE_DBSTRING ?= $(shell grep DB_URL .env | cut -d '=' -f2-)

# Auto generate .PHONY
.PHONY: $(MAKECMDGOALS)

##@ Helps
help: ## Show help message/commands
	@echo "$(GREEN)Makfile: Available commands for this project:$(NC)"
	@awk 'BEGIN {FS = ":.*##"} \
	/^[a-zA-Z_-]+:.*##/ { \
		printf "$(YELLOW)* $(GREEN)%-20s$(NC) %s\n", $$1 "$(NC):", $$2 \
	}' $(MAKEFILE_LIST) | sort

##@ Commands
dev: ## Start development server with live reload (Air)
	@echo "$(YELLOW)Starting development server...$(NC)"
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "$(RED)Air not found. Install with:$(NC)"; \
		echo "$(CYAN)go install github.com/air-verse/air@latest$(NC)"; \
		exit 1; \
	fi

start: ## Run the production binary (requires build first)
	@echo "$(YELLOW)Starting $(PROJECT_NAME)...$(NC)"
	@./$(BINARY_NAME)

build: ## Build optimized binary for current OS
	@echo "$(YELLOW)Building optimized $(PROJECT_NAME) for local OS...$(NC)"
	@go build -ldflags="-w -s" -trimpath -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)Build complete: $(BINARY_NAME)$(NC)"
	@ls -lh $(BINARY_NAME) | awk '{print "$(CYAN)Binary size: " $$5 "$(NC)"}'

format: ## Format all Go source code
	@echo "$(YELLOW)Formatting code...$(NC)"
	@go fmt ./...

deps: ## Download and cache all dependencies
	@echo "$(YELLOW)Downloading dependencies...$(NC)"
	@go mod download

tidy: ## Remove unused dependencies and clean go.mod
	@echo "$(YELLOW)Tidying dependencies...$(NC)"
	@go mod tidy
	@echo "$(GREEN)Dependencies cleaned$(NC)"

migrate-up: ## Run all pending migrations [GOOSE_DBSTRING=<db_url>]
	@echo "$(YELLOW)Running migrations...$(NC)"
	@if command -v goose >/dev/null 2>&1; then \
		GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING="$(GOOSE_DBSTRING)" goose -dir $(GOOSE_DIR) up; \
		echo "$(GREEN)Migrations applied$(NC)"; \
	else \
		echo "$(RED)goose not found. Install with:$(NC)"; \
		echo "$(CYAN)go install github.com/pressly/goose/v3/cmd/goose@latest$(NC)"; \
		exit 1; \
	fi

migrate-down: ## Rollback last migration [GOOSE_DBSTRING=<db_url>]
	@echo "$(YELLOW)Rolling back last migration...$(NC)"
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING="$(GOOSE_DBSTRING)" goose -dir $(GOOSE_DIR) down

migrate-status: ## Show migration status [GOOSE_DBSTRING=<db_url>]
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING="$(GOOSE_DBSTRING)" goose -dir $(GOOSE_DIR) status

migrate-reset: ## Rollback all migrations [GOOSE_DBSTRING=<db_url>]
	@echo "$(RED)Resetting all migrations...$(NC)"
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING="$(GOOSE_DBSTRING)" goose -dir $(GOOSE_DIR) reset

sqlc: ## Generate type-safe Go code from SQL queries
	@echo "$(YELLOW)Generating sqlc code...$(NC)"
	@if command -v sqlc >/dev/null 2>&1; then \
		cd sqlpg && sqlc generate; \
		echo "$(GREEN)sqlc generation complete$(NC)"; \
	else \
		echo "$(RED)sqlc not found. Install with:$(NC)"; \
		echo "$(CYAN)go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest$(NC)"; \
		exit 1; \
	fi
