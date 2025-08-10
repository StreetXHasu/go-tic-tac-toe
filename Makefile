.PHONY: build test test-coverage clean lint docker-build docker-up docker-down

# Variables
SERVICES := gateway services/auth services/user services/game services/chat
PKG := pkg

# Build all services
build:
	@echo "Building all services..."
	@for service in $(SERVICES); do \
		echo "Building $$service..."; \
		cd $$service && go build -o bin/$$(basename $$service) . && cd -; \
	done

# Run tests
test:
	@echo "Running tests..."
	@for service in $(SERVICES); do \
		echo "Testing $$service..."; \
		cd $$service && go test ./... && cd -; \
	done
	@echo "Testing pkg..."
	@cd $(PKG) && go test ./...

# Run BDD tests with Godog
test-bdd:
	@echo "Running BDD tests..."
	@for service in $(SERVICES); do \
		if [ -d "$$service/tests" ]; then \
			echo "Running BDD tests for $$service..."; \
			cd $$service && godog run tests/ && cd -; \
		fi; \
	done

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@for service in $(SERVICES); do \
		echo "Testing $$service with coverage..."; \
		cd $$service && go test -coverprofile=coverage.out ./... && cd -; \
	done
	@echo "Testing pkg with coverage..."
	@cd $(PKG) && go test -coverprofile=coverage.out ./...

# Clean
clean:
	@echo "Cleaning build artifacts..."
	@for service in $(SERVICES); do \
		rm -rf $$service/bin; \
		rm -f $$service/coverage.out; \
	done
	@rm -rf $(PKG)/bin
	@rm -f $(PKG)/coverage.out

# Linting
lint:
	@echo "Running linter..."
	@for service in $(SERVICES); do \
		echo "Linting $$service..."; \
		cd $$service && golangci-lint run && cd -; \
	done
	@echo "Linting pkg..."
	@cd $(PKG) && golangci-lint run

# Docker commands
docker-build:
	@echo "Building Docker images..."
	cd infra && docker-compose build

docker-up:
	@echo "Starting services..."
	cd infra && docker-compose up -d

docker-down:
	@echo "Stopping services..."
	cd infra && docker-compose down

docker-logs:
	@echo "Showing logs..."
	cd infra && docker-compose logs -f

# Log management commands
logs-frontend:
	@echo "Showing frontend logs..."
	cd infra && ./scripts/view-logs.sh frontend

logs-gateway:
	@echo "Showing gateway logs..."
	cd infra && ./scripts/view-logs.sh gateway

logs-auth:
	@echo "Showing auth service logs..."
	cd infra && ./scripts/view-logs.sh auth

logs-user:
	@echo "Showing user service logs..."
	cd infra && ./scripts/view-logs.sh user

logs-game:
	@echo "Showing game service logs..."
	cd infra && ./scripts/view-logs.sh game

logs-chat:
	@echo "Showing chat service logs..."
	cd infra && ./scripts/view-logs.sh chat

logs-postgres:
	@echo "Showing postgres logs..."
	cd infra && ./scripts/view-logs.sh postgres

logs-all:
	@echo "Showing all logs..."
	cd infra && ./scripts/view-logs.sh all

logs-error:
	@echo "Showing error logs..."
	cd infra && ./scripts/view-logs.sh error

logs-stats:
	@echo "Showing log statistics..."
	cd infra && ./scripts/view-logs.sh stats

logs-clear:
	@echo "Clearing all logs..."
	cd infra && ./scripts/view-logs.sh clear

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@for service in $(SERVICES); do \
		echo "Installing deps for $$service..."; \
		cd $$service && go mod tidy && cd -; \
	done
	@echo "Installing deps for pkg..."
	@cd $(PKG) && go mod tidy

# Install Godog
install-godog:
	@echo "Installing Godog..."
	go install github.com/cucumber/godog/cmd/godog@latest

# Install golangci-lint
install-lint:
	@echo "Installing golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Complete setup of development tools
setup: install-godog install-lint
	@echo "Setup complete!"

# Help
help:
	@echo "Available commands:"
	@echo "  build        - Build all services"
	@echo "  test         - Run unit tests"
	@echo "  test-bdd     - Run BDD tests with Godog"
	@echo "  test-coverage- Run tests with coverage"
	@echo "  clean        - Clean build artifacts"
	@echo "  lint         - Run linter"
	@echo "  deps         - Install dependencies"
	@echo "  setup        - Install development tools"
	@echo "  docker-build - Build Docker images"
	@echo "  docker-up    - Start services"
	@echo "  docker-down  - Stop services"
	@echo "  docker-logs  - Show Docker logs"
	@echo ""
	@echo "Logging commands:"
	@echo "  logs-frontend - Show frontend logs"
	@echo "  logs-gateway  - Show gateway logs"
	@echo "  logs-auth     - Show auth service logs"
	@echo "  logs-user     - Show user service logs"
	@echo "  logs-game     - Show game service logs"
	@echo "  logs-chat     - Show chat service logs"
	@echo "  logs-postgres - Show postgres logs"
	@echo "  logs-all      - Show all logs"
	@echo "  logs-error    - Show error logs only"
	@echo "  logs-stats    - Show log statistics"
	@echo "  logs-clear    - Clear all logs"
	@echo "  help          - Show this help" 