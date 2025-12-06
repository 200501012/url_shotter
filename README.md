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

## 🛣️ API Routes

### Health Check

**GET** `/health`

Checks if the application is working.

**Response:**

```
200 OK
Hello, World!, Olha como o mundo e' maravilhoso
```

**Example:**

```bash
curl http://localhost:8080/health
```

---

### Create URL

**POST** `/api/urls`

Creates a new shortened URL.

**Request Body:**

```json
{
  "original_url": "https://example.com"
}
```

**Response:**

```json
{
  "id": 1,
  "original_url": "https://example.com",
  "short_code": "a1b2c3d4",
  "click_count": 0,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

**Status Codes:**

- `201 Created` - URL created successfully
- `400 Bad Request` - Invalid URL
- `500 Internal Server Error` - Error creating URL

**Example:**

```bash
curl -X POST http://localhost:8080/api/urls \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://example.com"}'
```

---

### Get URL by ID

**GET** `/api/urls/:id`

Gets information about a URL by its ID or short code.

**Parameters:**

- `id` (path) - ID or short code of the URL

**Response:**

```json
{
  "id": 1,
  "original_url": "https://example.com",
  "short_code": "a1b2c3d4",
  "click_count": 5,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

**Status Codes:**

- `200 OK` - URL found
- `404 Not Found` - URL not found

**Example:**

```bash
# By ID
curl http://localhost:8080/api/urls/1

# By short code
curl http://localhost:8080/api/urls/a1b2c3d4
```

---

### Redirect to Original URL

**GET** `/:shortCode`

Redirects to the original URL using the short code. Automatically increments the click counter.

**Parameters:**

- `shortCode` (path) - Short code of the URL

**Response:**

```
301 Moved Permanently
Location: https://example.com
```

**Status Codes:**

- `301 Moved Permanently` - Redirect successful
- `404 Not Found` - URL not found

**Example:**

```bash
# Redirects to the original URL
curl -L http://localhost:8080/a1b2c3d4

# Or access directly in the browser
# http://localhost:8080/a1b2c3d4
```

---

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
