<p align="center">
  <img src="logo.svg" alt="Forge" width="480"/>
</p>

<p align="center">
  <strong>Scaffold production-ready fullstack projects in seconds.</strong>
</p>

<p align="center">
  <a href="#installation">Installation</a> ·
  <a href="#quick-start">Quick Start</a> ·
  <a href="#stacks">Stacks</a> ·
  <a href="#features">Features</a>
</p>

---

Forge is a CLI that generates fullstack project scaffolds from opinionated templates. Pick a stack, toggle the features you need, and get a ready-to-run monorepo — with dependencies installed and boilerplate wired up.

## Stacks

| Stack | Frontend | Backend | Use case |
|-------|----------|---------|----------|
| **`web`** | Next.js 14 (App Router) | Go + Echo | Fullstack web app with separate API |
| **`mobile`** | Expo / React Native | Go + Echo | Mobile app with separate API |
| **`web-fullstack`** | Next.js 14 (App Router) | — | Web app with Supabase direct (no Go backend) |

All stacks include **TypeScript**, **Tailwind CSS**, and **Supabase** out of the box.

## Features

Optional modules you can toggle during project creation:

| Feature | Description | Compatible stacks |
|---------|-------------|-------------------|
| **`auth`** | JWT auth flow — login, register, protected routes | `web` `mobile` `web-fullstack` |
| **`database`** | Supabase client, query builder, repository pattern | `web` `mobile` `web-fullstack` |
| **`api`** | Echo router, middleware, CRUD handlers | `web` `mobile` |

## Requirements

- Go 1.21+
- Git
- Node.js 18+ and npm (for frontend dependencies)

## Installation

```bash
git clone https://github.com/bigbytes/forge.git
cd forge
go build -o forge ./cmd/forge

# Optional: add to PATH
sudo mv forge /usr/local/bin/
```

## Quick Start

```bash
# Interactive — prompts for stack and features
forge new my-app

# Non-interactive
forge new my-app --stack web --features auth,database,api

# Web-only (no Go backend)
forge new my-app --stack web-fullstack --features auth,database

# Start developing
cd my-app
npm run dev
```

## Usage

```
forge new <project-name> [flags]
```

| Flag | Short | Description |
|------|-------|-------------|
| `--stack` | `-s` | Stack type: `web`, `mobile`, or `web-fullstack` |
| `--features` | | Comma-separated: `auth`, `database`, `api` |
| `--output` | `-o` | Output directory |
| `--force` | `-f` | Overwrite existing directory |
| `--skip-install` | | Skip automatic dependency installation |

```bash
forge version    # Show version
forge --help     # Show help
```

## Project Structure

### `web` / `mobile`

```
my-app/
├── frontend/  (or mobile/)   # Next.js or Expo app
│   ├── src/
│   │   ├── app/              # Pages / screens
│   │   ├── components/
│   │   └── lib/
│   └── package.json
├── backend/                   # Go API server
│   ├── cmd/server/
│   ├── internal/
│   └── go.mod
└── README.md
```

### `web-fullstack`

```
my-app/
├── src/
│   ├── app/                   # Next.js App Router
│   ├── components/
│   └── lib/                   # Supabase client & utils
├── package.json
└── README.md
```

## Environment Variables

### Backend (`.env`) — `web` and `mobile` stacks

```env
PORT=8080
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-anon-key
SUPABASE_SERVICE_KEY=your-service-key
JWT_SECRET=your-jwt-secret
```

### Frontend (`.env.local`)

```env
# web / mobile
NEXT_PUBLIC_API_URL=http://localhost:8080

# web-fullstack
NEXT_PUBLIC_SUPABASE_URL=https://your-project.supabase.co
NEXT_PUBLIC_SUPABASE_ANON_KEY=your-anon-key
```

## Project Name Rules

- Starts with a letter
- Letters, numbers, hyphens, underscores only
- Max 64 characters

## Development

```bash
# Build
go build -o forge ./cmd/forge

# Test
go test ./...

# Vet
go vet ./...
```

## License

MIT
