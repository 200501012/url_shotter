# URL Shortener

API for URL shortening built with Go and Fiber.

## 🚀 How to Run

### Option 1: Using Docker Compose (Recommended)

```bash
# Start all services (app + postgres + pgadmin)
docker-compose up -d

# View logs
docker-compose logs -f app

# Stop all services
docker-compose down
```

### Option 2: Using Docker Hub Image

If you did `docker pull elton971/url_shotter:latest`, you need:

1. **Have a PostgreSQL database running** (you can use docker-compose only for postgres)

2. **Run the image with environment variables:**

```bash
docker run -d \
  --name url_shotter \
  -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=5433 \
  -e DB_USER="eltoncavele8@gmail.com" \
  -e DB_PASSWORD="URLShortenerPass" \
  -e DB_NAME="URL_Shortener" \
  elton971/url_shotter:latest
```

**OR** if PostgreSQL is on the same Docker network:

```bash
# Create a network
docker network create url_shotter_net

# Run PostgreSQL (if not already running)
docker run -d \
  --name postgres_db \
  --network url_shotter_net \
  -e POSTGRES_USER="eltoncavele8@gmail.com" \
  -e POSTGRES_PASSWORD="URLShortenerPass" \
  -e POSTGRES_DB="URL_Shortener" \
  -p 5433:5432 \
  postgres:16-alpine

# Run the application on the same network
docker run -d \
  --name url_shotter \
  --network url_shotter_net \
  -p 8080:8080 \
  -e DB_HOST=postgres_db \
  -e DB_PORT=5432 \
  -e DB_USER="eltoncavele8@gmail.com" \
  -e DB_PASSWORD="URLShortenerPass" \
  -e DB_NAME="URL_Shortener" \
  elton971/url_shotter:latest
```

### Option 3: All-in-One Image (App + PostgreSQL)

The all-in-one image includes both the application and PostgreSQL in a single container:

```bash
# Pull the image
docker pull elton971/url_shotter:all-in-one

# Run with environment variables
docker run -d \
  --name url_shotter \
  -p 8080:8080 \
  -p 5432:5432 \
  -e DB_USER="eltoncavele8@gmail.com" \
  -e DB_PASSWORD="URLShortenerPass" \
  -e DB_NAME="URL_Shortener" \
  -v url_shotter_data:/var/lib/postgresql/data \
  elton971/url_shotter:all-in-one
```

**Note:** The all-in-one image automatically:

- Initializes PostgreSQL on first run
- Creates the database user and database
- Starts PostgreSQL in the background
- Launches the Go application

## 📝 Environment Variables

The application requires the following environment variables:

- `DB_HOST` - PostgreSQL host (default: localhost)
- `DB_PORT` - PostgreSQL port (default: 5432)
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name

## 🔍 Check if it's working

```bash
# Health check
curl http://localhost:8080/health
```

## 📦 Build Images

### Standard Image (App only)

```bash
# Local build
docker build -t elton971/url_shotter:latest .

# Push to Docker Hub
docker login
docker push elton971/url_shotter:latest
```

### All-in-One Image (App + PostgreSQL)

```bash
# Local build
docker build -f Dockerfile.all-in-one -t elton971/url_shotter:all-in-one .

# Push to Docker Hub
docker login
docker push elton971/url_shotter:all-in-one
```

## 🛠️ Development

```bash
# Run locally
go run ./cmd/app/main.go

# Build binary
go build -o url_shotter ./cmd/app/main.go

# Run binary
./url_shotter
```

## 📚 Project Structure

```
url_shotter/
├── cmd/
│   └── app/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   └── http/                # HTTP handlers and server
├── docker-compose.yml       # Docker Compose configuration
├── Dockerfile               # Standard Docker image
├── Dockerfile.all-in-one    # All-in-one Docker image
└── README.md               # This file
```

## 🐳 Docker Images

- `elton971/url_shotter:latest` - Application only (requires external PostgreSQL)
- `elton971/url_shotter:all-in-one` - Application + PostgreSQL (self-contained)

## 📄 License

MIT
