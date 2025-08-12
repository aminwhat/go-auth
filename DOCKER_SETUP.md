# Docker Setup Guide for Go Auth Service

This guide provides comprehensive instructions for running the Go Auth service using Docker and Docker Compose.

## 🚀 Quick Start

1. **Clone and navigate to the project**

   ```bash
   git clone <your-repo>
   cd go-auth
   ```

2. **Start the services**

   ```bash
   docker-compose up -d
   ```

3. **Access the application**
   - API: <http://localhost:3000>
   - Swagger Docs: <http://localhost:3000/swagger/index.html>
   - Health Check: <http://localhost:3000/health>
   - MongoDB: localhost:27017

## 📋 Prerequisites

- Docker Engine 20.10 or higher
- Docker Compose 2.0 or higher
- Available ports: 3000 (API) and 27017 (MongoDB)

## 🏗️ Architecture

The Docker setup includes two services:

### go-auth Service

- **Image**: Built from local Dockerfile
- **Port**: 3000:3000 (host:container)
- **Health Check**: HTTP GET /health every 30s
- **Dependencies**: MongoDB must be healthy before starting

### mongodb Service

- **Image**: mongo:8.0
- **Port**: 27017:27017 (host:container)
- **Credentials**: root/my-secret-pw
- **Database**: go-auth
- **Health Check**: MongoDB ping every 10s
- **Persistent Storage**: Docker volume `mongo_data`

## 🛠️ Management Commands

### Using Make (Recommended)

```bash
# Start services
make docker-run

# Stop services  
make docker-stop

# View logs
make docker-logs

# Clean up (removes containers and volumes)
make docker-clean

# Rebuild and restart
make docker-rebuild

# Open MongoDB shell
make mongo-shell

# Check service health
make docker-health
```

### Using Docker Compose Directly

```bash
# Build images
docker-compose build

# Start services in background
docker-compose up -d

# Start services with logs visible
docker-compose up

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v

# View logs
docker-compose logs -f

# View service status
docker-compose ps
```

## 🔧 Configuration

### Environment Variables

The application uses these environment variables (configured in docker-compose.yml):

| Variable | Value | Description |
|----------|-------|-------------|
| MONGO_URI | mongodb://mongodb:27017 | Internal MongoDB connection |
| MONGO_USERNAME | root | MongoDB username |
| MONGO_PASSWORD | my-secret-pw | MongoDB password |
| MONGO_DB_NAME | go-auth | Database name |

### Customization

To customize the configuration:

1. **Change ports**: Edit the `ports` section in `docker-compose.yml`
2. **Change credentials**: Update environment variables in `docker-compose.yml`
3. **Add volumes**: Mount additional directories as needed

Example customization:

```yaml
services:
  go-auth:
    ports:
      - "8080:3000"  # Change API port to 8080
    environment:
      MONGO_PASSWORD: your-secure-password  # Change password
```

## 🧪 Testing

### Automated Testing

Run the comprehensive test script:

```bash
chmod +x docker-test.sh
./docker-test.sh
```

This script tests:

- Docker installation
- Image building
- Service startup
- Health checks
- API endpoints
- Database connectivity

### Manual Testing

1. **Health Check**

   ```bash
   curl http://localhost:3000/health
   ```

2. **API Registration**

   ```bash
   curl -X POST http://localhost:3000/auth/signup \
     -H "Content-Type: application/json" \
     -d '{"phoneNumber": "09123456789"}'
   ```

3. **Swagger Documentation**
   Open: <http://localhost:3000/swagger/index.html>

## 📊 Monitoring

### View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f go-auth
docker-compose logs -f mongodb

# Last N lines
docker-compose logs --tail=50 go-auth
```

### Service Health

```bash
# Check container status
docker-compose ps

# Check health status
docker-compose ps --format "table {{.Name}}\t{{.Status}}\t{{.Ports}}"

# Monitor resource usage
docker stats $(docker-compose ps -q)
```

## 🗄️ Database Management

### Access MongoDB Shell

```bash
# Using Make
make mongo-shell

# Using Docker Compose
docker-compose exec mongodb mongosh -u root -p my-secret-pw --authenticationDatabase admin go-auth
```

### Common Database Operations

```javascript
// List collections
show collections

// Find all users
db.users.find().pretty()

// Count users
db.users.countDocuments()

// Find user by phone
db.users.find({"phoneNumber": "09123456789"})

// Check indexes
db.users.getIndexes()
```

### Backup and Restore

```bash
# Backup
make mongo-backup

# Restore
make mongo-restore

# Manual backup
docker-compose exec mongodb mongodump -u root -p my-secret-pw --authenticationDatabase admin --db go-auth --out /tmp/backup

# Manual restore
docker-compose exec mongodb mongorestore -u root -p my-secret-pw --authenticationDatabase admin --db go-auth /tmp/backup/go-auth
```

## 🚨 Troubleshooting

### Common Issues

#### 1. Port Already in Use

**Error**: `bind: address already in use`

**Solution**:

- Check what's using the port: `lsof -i :3000`
- Change ports in docker-compose.yml
- Stop conflicting services

#### 2. MongoDB Connection Failed

**Error**: `mongo ping error`

**Solutions**:

- Check MongoDB container health: `docker-compose ps`
- View MongoDB logs: `docker-compose logs mongodb`
- Restart services: `docker-compose restart`

#### 3. Build Failures

**Error**: Various build-related errors

**Solutions**:

- Clean Docker cache: `docker system prune -f`
- Rebuild images: `docker-compose build --no-cache`
- Check Dockerfile syntax and dependencies

#### 4. Health Check Failing

**Error**: Service marked as unhealthy

**Solutions**:

- Check application logs: `docker-compose logs go-auth`
- Test endpoints manually: `curl http://localhost:3000/health`
- Verify environment variables are correct

### Debug Commands

```bash
# Enter container shell
docker-compose exec go-auth /bin/sh

# Check environment variables
docker-compose exec go-auth env

# Test internal connectivity
docker-compose exec go-auth wget -qO- http://localhost:3000/health

# Check network connectivity
docker-compose exec go-auth ping mongodb
```

## 🔒 Security Considerations

### Development vs Production

This setup is optimized for development. For production:

1. **Change default credentials**

   ```yaml
   environment:
     MONGO_USERNAME: secure_username
     MONGO_PASSWORD: very_secure_password_123!
     JWT_SECRET_KEY: very_long_random_secret_key
   ```

2. **Use secrets management**

   ```yaml
   secrets:
     mongo_password:
       file: ./secrets/mongo_password.txt
   ```

3. **Limit network exposure**

   ```yaml
   ports:
     - "127.0.0.1:3000:3000"  # Bind to localhost only
   ```

4. **Enable SSL/TLS**
   - Configure reverse proxy (nginx/traefik)
   - Use MongoDB with SSL
   - Set up proper certificates

## 📈 Production Deployment

### Docker Compose Override

Create `docker-compose.prod.yml`:

```yaml
services:
  go-auth:
    environment:
      GIN_MODE: release
      JWT_SECRET_KEY: ${JWT_SECRET_KEY}
    restart: unless-stopped
    
  mongodb:
    environment:
      MONGO_INITDB_ROOT_PASSWORD: ${MONGO_ROOT_PASSWORD}
    volumes:
      - /var/lib/mongodb:/data/db
```

Run with:

```bash
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

### Orchestration

For production, consider using:

- **Kubernetes** with Helm charts
- **Docker Swarm** for clustering
- **Portainer** for management UI
- **Traefik** for reverse proxy and load balancing

## 🔄 Updates and Maintenance

### Updating the Application

```bash
# Pull latest code
git pull

# Rebuild and restart
make docker-rebuild

# Or manually
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

### Database Migrations

When schema changes are needed:

1. Stop the application: `docker-compose stop go-auth`
2. Run migration scripts via MongoDB shell
3. Restart: `docker-compose start go-auth`

### Cleanup

```bash
# Remove unused containers and images
docker system prune

# Remove specific project containers
docker-compose down --rmi all --volumes --remove-orphans

# Reset everything
make docker-clean
docker system prune -a
```

## 📚 Additional Resources

- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [MongoDB Docker Hub](https://hub.docker.com/_/mongo)
- [Go Docker Best Practices](https://docs.docker.com/language/golang/)
- [Production Docker Guide](https://docs.docker.com/config/containers/start-containers-automatically/)

---

For questions or issues, check the logs first:

```bash
docker-compose logs -f
```
