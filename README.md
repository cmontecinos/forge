# Forge

CLI tool to create projects with your preferred stacks (Next.js/Expo + Go + Supabase).

## Build

```bash
go build ./cmd/forge
```

## Usage

```bash
# Create a new project (interactive)
./forge new my-project

# Create with specific stack
./forge new my-project --stack web

# Create with features
./forge new my-project --stack mobile --features auth,database,api
```

## Stacks

- **Web**: Next.js (App Router + Tailwind + TypeScript) + Go (Echo) + Supabase
- **Mobile**: Expo (React Native + TypeScript) + Go (Echo) + Supabase

## Features

- **Auth**: Login/register via backend, JWT validation
- **Database**: Go-Supabase connection, base models
- **API**: Echo router, middlewares, example handlers
