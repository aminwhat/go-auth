# Go Auth Makefile

.PHONY: help build run stop clean logs shell test swagger docker-build docker-run docker-stop docker-clean

# Default target
help:
	@echo "Available commands:"
	@echo "  build         - Build the Go application"
	@echo "  run           - Run the application locally"
	@echo "  test          - Run tests"
	@echo "  swagger       - Generate Swagger documentation"
	@echo "  clean         - Clean build artifacts"
	@echo ""
	@echo "Docker commands:"
	@echo "  docker-build  - Build Docker images"
	@echo "  docker-run    - Start services with docker-compose"
	@echo "  docker-stop   - Stop services"
	@echo "  docker-clean  - Clean Docker containers and volumes"
	@echo "  docker-logs   - Show logs from services"
	@echo "  docker-shell  - Open shell in go-auth container"
	@echo ""
	@echo "Development:"
	@echo "  dev-setup     - Set up development environment"
	@echo "  dev-run       - Run in development mode with hot reload"

# Local development
build:
	go build -o bin/main .

run:
	go run main.go

test:
	go test ./...

swagger:
	swag init

clean:
	rm -rf bin/
	go clean

# Docker commands
docker-build:
	docker-compose build

docker-run:
	docker-compose up -d

docker-stop:
	docker-compose down

docker-clean:
	docker-compose down -v --remove-orphans
	docker system prune -f

docker-logs:
	docker-compose logs -f

docker-shell:
	docker-compose exec go-auth /bin/sh

# Combined commands
docker-restart: docker-stop docker-run

docker-rebuild: docker-stop docker-build docker-run

# Development setup
dev-setup:
	@echo "Setting up development environment..."
	go mod download
	@if [ ! -f .env ]; then \
		echo "Creating .env file from .env.docker template..."; \
		cp .env.docker .env; \
		echo "Please edit .env file with your configuration"; \
	fi

dev-run: swagger
	@echo "Starting development server..."
	go run main.go

# MongoDB commands (for Docker setup)
mongo-shell:
	docker-compose exec mongodb mongosh -u root -p my-secret-pw --authenticationDatabase admin go-auth

mongo-backup:
	@echo "Creating MongoDB backup..."
	docker-compose exec mongodb mongodump -u root -p my-secret-pw --authenticationDatabase admin --db go-auth --out /data/backup

mongo-restore:
	@echo "Restoring MongoDB backup..."
	docker-compose exec mongodb mongorestore -u root -p my-secret-pw --authenticationDatabase admin --db go-auth /data/backup/go-auth

# Production commands
prod-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s' -o bin/main .

prod-run:
	GIN_MODE=release ./bin/main

# Health check
health:
	@curl -s http://localhost:3000/swagger/index.html > /dev/null && echo "✅ Service is running" || echo "❌ Service is not responding"

# Docker health check
docker-health:
	@docker-compose ps
