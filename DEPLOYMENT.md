# Sentinel Auth - Deployment Guide

This guide provides instructions for deploying the Sentinel Auth system using Docker Compose for both development and production environments.

## Quick Start

### Development Deployment

```bash
# Clone the repository
git clone <repository-url>
cd sentinel

# Start development environment
docker compose -f docker-compose.dev.yml up -d

# View logs
docker compose -f docker-compose.dev.yml logs -f
```

The development environment includes:
- **Auth Service**: `http://localhost:8080`
- **Database**: PostgreSQL on `localhost:5432`
- **Database Admin**: Adminer on `http://localhost:8081`

### Production Deployment

```bash
# Clone the repository
git clone <repository-url>
cd sentinel

# Copy and configure production environment
cp .env.prod .env.prod.local
# Edit .env.prod.local with your secure values

# Start production environment
docker-compose -f docker-compose.prod.yml --env-file .env.prod.local up -d

# View logs
docker-compose -f docker-compose.prod.yml logs -f
```

## Environment Configuration

### Development Environment
- Uses `.env.dev` for configuration
- Includes database admin interface (Adminer)
- Simplified security settings
- Volume mounting for development

### Production Environment
- Uses `.env.prod` (copy to `.env.prod.local` and customize)
- **IMPORTANT**: Change default passwords and secrets
- Includes Nginx reverse proxy with rate limiting
- Resource limits and health checks
- SSL-ready configuration

## Required Environment Variables

### Application Variables
```env
PORT=8080
ENVIRONMENT=production
JWT_SECRET=your-secure-jwt-secret-here
DB_HOST=postgres
DB_PORT=5432
DB_NAME=sentinel_prod
DB_USER=sentinel_user
DB_PASSWORD=your-secure-db-password-here
DB_SSLMODE=require
```

### PostgreSQL Variables
```env
POSTGRES_DB=sentinel_prod
POSTGRES_USER=sentinel_user
POSTGRES_PASSWORD=your-secure-db-password-here
```

## API Endpoints

### Health Check
- `GET /api/v1/health` - Service health status

### Authentication (Planned)
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/refresh` - Token refresh
- `POST /api/v1/auth/logout` - User logout
- `GET /api/v1/auth/oauth/callback` - OAuth callback

## Database Schema

The system automatically creates the following tables:
- `users` - User accounts
- `sessions` - Session management
- `oauth_providers` - OAuth provider configurations
- `oauth_accounts` - User OAuth account links
- `password_reset_tokens` - Password reset functionality
- `email_verification_tokens` - Email verification

## Production Security

### Before Production Deployment:

1. **Change Default Credentials**
   ```bash
   # Generate secure JWT secret
   openssl rand -hex 32
   
   # Generate secure database password
   openssl rand -hex 16
   ```

2. **SSL/TLS Setup** (Recommended)
   - Obtain SSL certificates
   - Place certificates in `nginx/ssl/`
   - Uncomment HTTPS configuration in `nginx/nginx.conf`

3. **Rate Limiting**
   - Login endpoints: 5 requests per minute
   - General API: 10 requests per second
   - Configurable in `nginx/nginx.conf`

4. **Resource Limits**
   - App: 512MB memory limit, 0.5 CPU
   - Database: 1GB memory limit, 1.0 CPU

## Maintenance

### Backup Database
```bash
docker-compose -f docker-compose.prod.yml exec postgres pg_dump -U sentinel_user sentinel_prod > backup.sql
```

### Restore Database
```bash
docker-compose -f docker-compose.prod.yml exec -T postgres psql -U sentinel_user sentinel_prod < backup.sql
```

### View Logs
```bash
# All services
docker-compose -f docker-compose.prod.yml logs -f

# Specific service
docker-compose -f docker-compose.prod.yml logs -f sentinel-auth
```

### Update Application
```bash
# Pull latest code
git pull

# Rebuild and restart
docker-compose -f docker-compose.prod.yml up -d --build
```

## Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Check if PostgreSQL container is healthy: `docker-compose ps`
   - Verify environment variables match between services

2. **Port Already in Use**
   - Change ports in docker-compose files
   - Stop conflicting services: `sudo lsof -i :8080`

3. **Permission Denied**
   - Ensure Docker daemon is running
   - Add user to docker group: `sudo usermod -aG docker $USER`

### Health Checks
- Application: `curl http://localhost:8080/api/v1/health`
- Database: `docker-compose exec postgres pg_isready -U sentinel_user`

## Monitoring

The production setup includes:
- Application health checks every 30 seconds
- Database health checks every 10 seconds
- Automatic service restart on failure
- Resource usage monitoring via Docker stats