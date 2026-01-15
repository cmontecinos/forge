# {{.ProjectName}}

A web application built with Next.js frontend and Go backend, using Supabase for the database.

## Project Structure

```
{{.ProjectName}}/
├── frontend/          # Next.js App Router + Tailwind CSS
│   ├── src/
│   │   ├── app/       # App Router pages
│   │   └── lib/       # Utilities and API client
│   └── package.json
├── backend/           # Go Echo server
│   ├── cmd/server/    # Entry point
│   └── internal/      # Private packages
│       ├── config/    # Configuration
│       └── server/    # HTTP server
└── package.json       # Root workspace config
```

## Getting Started

### Prerequisites

- Node.js 18+
- Go 1.21+
- Supabase account (for production)

### Setup

1. Install frontend dependencies:
   ```bash
   npm run install:frontend
   ```

2. Copy environment files:
   ```bash
   cp backend/.env.example backend/.env
   cp frontend/.env.example frontend/.env.local
   ```

3. Configure environment variables in both `.env` files.

### Development

Run both frontend and backend:
```bash
npm run dev
```

Or run them separately:
```bash
# Terminal 1 - Backend
npm run dev:backend

# Terminal 2 - Frontend
npm run dev:frontend
```

- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

### API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /health | Health check |
| * | /api/v1/* | API routes (add your endpoints here) |

## Architecture

```
Browser → Next.js (port 3000) → Go API (port 8080) → Supabase
```

The frontend never connects directly to Supabase. All data flows through the Go backend, which:
- Validates requests
- Handles authentication
- Manages database operations
- Returns JSON responses

## Environment Variables

### Backend (.env)

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Server port | 8080 |
| SUPABASE_URL | Supabase project URL | - |
| SUPABASE_KEY | Supabase anon/service key | - |

### Frontend (.env.local)

| Variable | Description | Default |
|----------|-------------|---------|
| NEXT_PUBLIC_API_URL | Backend API URL | http://localhost:8080 |

## Building for Production

```bash
# Build frontend
npm run build

# Build backend
cd backend && go build -o bin/server ./cmd/server
```

## Stack

- **Frontend**: Next.js 14, React, TypeScript, Tailwind CSS
- **Backend**: Go, Echo, godotenv
- **Database**: Supabase (PostgreSQL)
