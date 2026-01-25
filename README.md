# Forge

CLI tool to create full-stack projects with your preferred stacks. Generate production-ready monorepos with Next.js or Expo frontend, Go backend, and Supabase integration.

## Requirements

- Go 1.21+
- Git
- Node.js 18+ and npm (for frontend dependencies)
- Supabase account (for database and auth features)

## Installation

```bash
# Clone the repository
git clone https://github.com/bigbytes/forge.git
cd forge

# Build the CLI
go build -o forge ./cmd/forge

# Optional: Move to PATH
sudo mv forge /usr/local/bin/
```

## Quick Start

```bash
# Interactive mode - prompts for stack and features
forge new my-app

# Follow the prompts to select your stack and features
# Dependencies are installed automatically
cd my-app
npm run dev  # For web projects
# or
npx expo start  # For mobile projects
```

## Usage

### Create a New Project

```bash
forge new <project-name> [flags]
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--stack` | `-s` | Stack type: `web` or `mobile` |
| `--features` | | Features to include (comma-separated): `auth`, `database`, `api` |
| `--output` | `-o` | Output directory for the new project |
| `--force` | `-f` | Overwrite existing directory |
| `--skip-install` | | Skip automatic dependency installation |

### Examples

```bash
# Interactive mode - prompts for stack and features
forge new my-app

# Specify stack, prompt for features
forge new my-app --stack web

# Fully non-interactive with all features
forge new my-app --stack mobile --features auth,database,api

# Custom output directory
forge new my-app -o /path/to/projects

# Overwrite existing directory
forge new my-app --force

# Skip dependency installation (faster, install manually later)
forge new my-app --skip-install
```

### Other Commands

```bash
# Show version
forge version

# Show help
forge --help
forge new --help
```

## Stacks

### Web Stack

Full-stack web application with Next.js frontend and Go backend.

**Structure:**
```
my-app/
├── frontend/           # Next.js application
│   ├── src/
│   │   ├── app/       # App Router pages
│   │   ├── components/
│   │   └── lib/       # Utilities and API client
│   ├── package.json
│   └── tailwind.config.js
├── backend/            # Go API server
│   ├── cmd/server/
│   ├── internal/
│   │   ├── server/    # HTTP handlers and routes
│   │   ├── supabase/  # Supabase client
│   │   ├── models/    # Data models
│   │   └── repository/# Data access layer
│   └── go.mod
└── README.md
```

**Tech Stack:**
- Next.js 14 (App Router)
- TypeScript
- Tailwind CSS
- Go with Echo framework
- Supabase (PostgreSQL + Auth)

### Mobile Stack

Mobile application with Expo frontend and Go backend.

**Structure:**
```
my-app/
├── mobile/             # Expo application
│   ├── src/
│   │   ├── app/       # Expo Router screens
│   │   ├── components/
│   │   └── lib/       # Utilities and API client
│   └── package.json
├── backend/            # Go API server (same as web)
│   └── ...
└── README.md
```

**Tech Stack:**
- Expo SDK (managed workflow)
- React Native
- TypeScript
- Expo Router
- Go with Echo framework
- Supabase (PostgreSQL + Auth)

## Features

Features are optional modules that add functionality to your project. Select the ones you need during project creation.

### Auth

Complete authentication system with JWT-based auth flow.

**Includes:**
- Backend: Login, register, refresh token, logout endpoints
- Backend: JWT validation middleware for protected routes
- Frontend: Auth context/provider with hooks
- Frontend: Login and register forms
- Frontend: Protected route handling

**Usage:**
```typescript
// Frontend - use the auth hook
const { user, login, logout, isAuthenticated } = useAuth();

// Login
await login(email, password);

// Access protected API
// Authorization header is automatically added
```

### Database

Supabase database integration with query builder and repository pattern.

**Includes:**
- Backend: Supabase client configuration
- Backend: Fluent query builder for database operations
- Backend: Example Item model with CRUD repository
- Backend: Repository pattern for data access

**Usage:**
```go
// Query builder
results, err := db.From("items").
    Select("*").
    Eq("user_id", userID).
    Order("created_at", false).
    Execute()

// Repository
repo := repository.NewItemRepository(supabaseClient)
items, err := repo.GetByUserID(ctx, userID)
```

### API

Structured API layer with Echo framework, middleware, and example handlers.

**Includes:**
- Backend: Router with route groups
- Backend: Standard middleware (CORS, logging, recovery)
- Backend: Consistent error response helpers
- Backend: Example CRUD handlers for items

**Usage:**
```go
// Error responses
return BadRequest(c, "Invalid input")
return NotFound(c, "Item not found")
return ValidationError(c, map[string]string{"title": "required"})

// Success responses
return c.JSON(http.StatusOK, items)
```

## Environment Variables

After creating a project, configure your environment:

### Backend (.env)

```env
PORT=8080
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-anon-key
SUPABASE_SERVICE_KEY=your-service-key
JWT_SECRET=your-jwt-secret
```

### Frontend (.env.local)

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
# or for mobile:
EXPO_PUBLIC_API_URL=http://localhost:8080
```

## Development

### Running the Backend

```bash
cd backend
go run ./cmd/server
```

### Running the Frontend (Web)

```bash
cd frontend
npm run dev
```

### Running the Frontend (Mobile)

```bash
cd mobile
npx expo start
```

## Project Name Rules

Project names must:
- Start with a letter
- Contain only letters, numbers, hyphens, and underscores
- Be at most 64 characters

Valid examples: `my-app`, `myApp123`, `my_project`
Invalid examples: `123app`, `-myapp`, `my app`, `my@app`

## License

MIT
