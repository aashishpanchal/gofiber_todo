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
GO_VERSION := >=1.26.0
BINARY_NAME :=./cmd/main.bin
MAIN_PATH := ./cmd/main.go

# Auto generate .PHONY
.PHONY: $(MAKECMDGOALS)

##@ Helps
help: ## - Show help message
	@echo "$(CYAN)$(PROJECT_NAME) Management Commands$(NC)"
	@awk 'BEGIN {FS = ":.*##"} \
	/^[a-zA-Z_-]+:.*##/ { printf "  $(CYAN)%-10s$(NC) %s\n", $$1, $$2 } \
	/^##@/ { printf "\n$(GREEN)%s$(NC)\n", substr($$0, 5) } \
	END { printf "\n$(GREEN)Usage:$(NC)\n  make $(CYAN)<target>$(NC)\n\n" }' $(MAKEFILE_LIST)

##@ Commands
dev: ## - Start development server with live reload (Air)
	@echo "$(YELLOW)Starting development server...$(NC)"
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "$(RED)Air not found. Install with:$(NC)"; \
		echo "$(CYAN)go install github.com/air-verse/air@latest$(NC)"; \
		exit 1; \
	fi

start: ## - Run the production binary (requires build first)
	@echo "$(YELLOW)Starting $(PROJECT_NAME)...$(NC)"
	@./$(BINARY_NAME)

build: ## - Build optimized binary for current OS
	@echo "$(YELLOW)Building optimized $(PROJECT_NAME) for local OS...$(NC)"
	@go build -ldflags="-w -s" -trimpath -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)Build complete: $(BINARY_NAME)$(NC)"
	@ls -lh $(BINARY_NAME) | awk '{print "$(CYAN)Binary size: " $$5 "$(NC)"}'

format: ## - Format all Go source code
	@echo "$(YELLOW)Formatting code...$(NC)"
	@go fmt ./...

deps: ## - Download and cache all dependencies
	@echo "$(YELLOW)Downloading dependencies...$(NC)"
	@go mod download

tidy: ## - Remove unused dependencies and clean go.mod
	@echo "$(YELLOW)Tidying dependencies...$(NC)"
	@go mod tidy
	@echo "$(GREEN)Dependencies cleaned$(NC)"
