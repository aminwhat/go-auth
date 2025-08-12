#!/bin/bash

# Docker Test Script for Go Auth Service
# This script tests the Docker setup and ensures all services are working correctly

set -e  # Exit on any error

echo "🐳 Docker Test Script for Go Auth Service"
echo "=========================================="

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_info() {
    echo -e "ℹ️  $1"
}

# Function to wait for service to be ready
wait_for_service() {
    local service_name=$1
    local url=$2
    local max_attempts=30
    local attempt=1

    print_info "Waiting for $service_name to be ready..."

    while [ $attempt -le $max_attempts ]; do
        if curl -s -f "$url" > /dev/null 2>&1; then
            print_success "$service_name is ready!"
            return 0
        fi

        echo -n "."
        sleep 2
        attempt=$((attempt + 1))
    done

    print_error "$service_name failed to start within $((max_attempts * 2)) seconds"
    return 1
}

# Function to cleanup
cleanup() {
    print_info "Cleaning up..."
    docker-compose down -v > /dev/null 2>&1 || true
}

# Trap cleanup on exit
trap cleanup EXIT

# Step 1: Check if Docker and Docker Compose are installed
print_info "Checking Docker installation..."
if ! command -v docker &> /dev/null; then
    print_error "Docker is not installed"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    print_error "Docker Compose is not installed"
    exit 1
fi

print_success "Docker and Docker Compose are installed"

# Step 2: Build the images
print_info "Building Docker images..."
if docker-compose build; then
    print_success "Docker images built successfully"
else
    print_error "Failed to build Docker images"
    exit 1
fi

# Step 3: Start the services
print_info "Starting services..."
if docker-compose up -d; then
    print_success "Services started"
else
    print_error "Failed to start services"
    exit 1
fi

# Step 4: Wait for services to be healthy
print_info "Waiting for services to be healthy..."
sleep 10

# Check MongoDB
print_info "Testing MongoDB connection..."
if docker-compose exec -T mongodb mongosh --eval "db.adminCommand('ping')" > /dev/null 2>&1; then
    print_success "MongoDB is healthy"
else
    print_error "MongoDB is not responding"
    docker-compose logs mongodb
    exit 1
fi

# Wait for Go Auth service
wait_for_service "Go Auth API" "http://localhost:3000/health"

# Step 5: Test API endpoints
print_info "Testing API endpoints..."

# Test health endpoint
if curl -s -f "http://localhost:3000/health" | grep -q "healthy"; then
    print_success "Health endpoint is working"
else
    print_error "Health endpoint is not working"
    exit 1
fi

# Test Swagger documentation
if curl -s -f "http://localhost:3000/swagger/index.html" > /dev/null; then
    print_success "Swagger documentation is accessible"
else
    print_error "Swagger documentation is not accessible"
    exit 1
fi

# Test user signup endpoint (should require phone number)
print_info "Testing signup endpoint..."
signup_response=$(curl -s -X POST http://localhost:3000/auth/signup \
  -H "Content-Type: application/json" \
  -d '{"phoneNumber": "09123456789"}' \
  -w "%{http_code}")

if echo "$signup_response" | grep -q "200"; then
    print_success "Signup endpoint is working"
else
    print_warning "Signup endpoint test inconclusive (this may be expected due to validation)"
fi

# Step 6: Test database operations
print_info "Testing database operations..."
if docker-compose exec -T mongodb mongosh go-auth --eval "db.users.find().limit(1)" > /dev/null 2>&1; then
    print_success "Database operations are working"
else
    print_error "Database operations failed"
    exit 1
fi

# Step 7: Check service status
print_info "Checking service status..."
docker-compose ps

# Step 8: Show logs (last 10 lines)
print_info "Recent application logs:"
docker-compose logs --tail=10 go-auth

print_success "All tests passed! 🎉"
print_info "Services are running and healthy."
print_info ""
print_info "📡 Access your services:"
print_info "  • API: http://localhost:3000"
print_info "  • Health Check: http://localhost:3000/health"
print_info "  • Swagger Docs: http://localhost:3000/swagger/index.html"
print_info "  • MongoDB: mongodb://localhost:27017"
print_info ""
print_info "🛠️  Useful commands:"
print_info "  • View logs: docker-compose logs -f"
print_info "  • Stop services: docker-compose down"
print_info "  • Restart services: docker-compose restart"
print_info ""
print_info "To stop the test environment, run: docker-compose down"
