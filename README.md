# {{.ProjectName}}

A web application built with Go backend and Supabase.

## Project Structure

```
{{.ProjectName}}/
├── backend/           # Go backend (Echo framework)
│   ├── cmd/server/    # Application entry point
│   ├── internal/      # Private application code
│   │   ├── config/    # Configuration management
│   │   └── server/    # HTTP server and routes
│   └── go.mod         # Go module definition
└── README.md          # This file
```

## Getting Started

### Prerequisites

- Go 1.21 or later
- Supabase account (optional, for database features)

### Backend Development

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Copy environment configuration:
   ```bash
   cp .env.example .env
   ```

3. Edit `.env` with your Supabase credentials (if using)

4. Run the server:
   ```bash
   go run ./cmd/server
   ```

5. Test the health endpoint:
   ```bash
   curl http://localhost:8080/health
   ```

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /api/v1/*` - API v1 routes (add your routes in `internal/server/routes.go`)

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Server port | 8080 |
| SUPABASE_URL | Supabase project URL | - |
| SUPABASE_KEY | Supabase anon/service key | - |
