# Go Auth

A Sample Project using Gin Framework that implements an auth-based system using JWT Token with OTP authentication and user management features.

## Features

- OTP-based registration and login
- JWT token authentication
- Rate limiting (3 OTP requests per phone number within 10 minutes)
- User management with pagination and search
- MongoDB database integration
- Swagger API documentation

## API Endpoints

### Authentication

- `POST /auth/signup` - Register user with phone number (generates OTP)
- `POST /auth/signup/confirm-otp` - Confirm OTP and get JWT token

### User Management (Requires Authentication)

- `GET /user/` - Get current user details
- `GET /user/{userId}` - Get user by ID
- `GET /user/all` - Get all users with pagination and phone search
  - Query parameters: `page`, `pageSize`, `phone`

## Commands

### Run the application

```sh
go run main.go
```

### Build the application

```sh
go build
```

### Swagger generate

```sh
swag init
```

## API Documentation

Once the application is running, access the Swagger documentation at:

```
http://localhost:3000/swagger/index.html
```

For detailed API examples and usage, see [API_EXAMPLES.md](./API_EXAMPLES.md)

## Docker Setup

### Prerequisites

- Docker and Docker Compose installed on your system

### Quick Start with Docker

1. **Clone the repository**

```sh
git clone <your-repo-url>
cd go-auth
```

2. **Set up environment file**

```sh
cp .env.docker .env
# Edit .env file if needed
```

3. **Build and run with Docker Compose**

```sh
docker-compose up -d
```

4. **Access the application**

- API: <http://localhost:3000>
- Swagger Documentation: <http://localhost:3000/swagger/index.html>
- MongoDB: localhost:27017

### Docker Commands

#### Using Make (Recommended)

```sh
# Build Docker images
make docker-build

# Start services
make docker-run

# Stop services
make docker-stop

# View logs
make docker-logs

# Clean up (removes containers and volumes)
make docker-clean

# Restart services
make docker-restart

# Rebuild and restart
make docker-rebuild
```

#### Using Docker Compose Directly

```sh
# Build images
docker-compose build

# Start services in background
docker-compose up -d

# Start services with logs
docker-compose up

# Stop services
docker-compose down

# Stop services and remove volumes
docker-compose down -v

# View logs
docker-compose logs -f

# View logs for specific service
docker-compose logs -f go-auth
```

### Services

The Docker setup includes:

- **go-auth**: The main Go application (Port: 3000)
- **mongodb**: MongoDB database (Port: 27017)

### Environment Configuration

The application uses these environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| MONGO_URI | mongodb://mongodb:27017 | MongoDB connection URI (internal) |
| MONGO_USERNAME | root | MongoDB username |
| MONGO_PASSWORD | my-secret-pw | MongoDB password |
| MONGO_DB_NAME | go-auth | Database name |
| JWT_SECRET_KEY | my_very_secret_key | JWT signing secret |
| GIN_MODE | release | Gin framework mode |
| PORT | 3000 | Server port |

### Database Management

#### MongoDB Shell Access

```sh
# Using Make
make mongo-shell

# Using Docker Compose
docker-compose exec mongodb mongosh -u root -p my-secret-pw --authenticationDatabase admin go-auth
```

#### Database Backup/Restore

```sh
# Backup
make mongo-backup

# Restore
make mongo-restore
```

### Development with Docker

#### Access Container Shell

```sh
# Using Make
make docker-shell

# Using Docker Compose
docker-compose exec go-auth /bin/sh
```

#### View Container Logs

```sh
# All services
make docker-logs

# Specific service
docker-compose logs -f go-auth
docker-compose logs -f mongodb
```

### Troubleshooting

#### Service Health Check

```sh
# Check if services are running
make docker-health

# Manual health check
curl http://localhost:3000/swagger/index.html
```

#### Common Issues

1. **Port already in use**
   - Change ports in docker-compose.yml
   - Stop conflicting services

2. **MongoDB connection issues**
   - Ensure MongoDB container is healthy: `docker-compose ps`
   - Check logs: `docker-compose logs mongodb`

3. **Build failures**
   - Clean Docker cache: `docker system prune -f`
   - Rebuild: `make docker-rebuild`

### Production Deployment

For production deployment, update the environment variables in docker-compose.yml:

```yaml
environment:
  MONGO_URI: mongodb://your-mongo-host:27017
  MONGO_USERNAME: your-username
  MONGO_PASSWORD: your-secure-password
  MONGO_DB_NAME: go-auth-prod
  JWT_SECRET_KEY: your-very-secure-jwt-secret
  GIN_MODE: release
```

## Local Development
