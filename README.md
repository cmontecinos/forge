<p align="center">
  <img src="logo.svg" alt="Forge" width="520"/>
</p>

<p align="center">
  <strong>Scaffold production-ready fullstack projects in seconds.</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go 1.21+"/>
  <img src="https://img.shields.io/badge/Next.js-14-000000?style=flat-square&logo=next.js&logoColor=white" alt="Next.js 14"/>
  <img src="https://img.shields.io/badge/Expo-React_Native-4630EB?style=flat-square&logo=expo&logoColor=white" alt="Expo"/>
  <img src="https://img.shields.io/badge/Supabase-Postgres+Auth-3FCF8E?style=flat-square&logo=supabase&logoColor=white" alt="Supabase"/>
  <img src="https://img.shields.io/badge/license-MIT-blue?style=flat-square" alt="MIT License"/>
</p>

<p align="center">
  <a href="#installation">Installation</a> &nbsp;&middot;&nbsp;
  <a href="#quick-start">Quick Start</a> &nbsp;&middot;&nbsp;
  <a href="#stacks">Stacks</a> &nbsp;&middot;&nbsp;
  <a href="#features">Features</a> &nbsp;&middot;&nbsp;
  <a href="#project-structure">Structure</a>
</p>

---

Forge is a CLI that generates fullstack project scaffolds from opinionated templates. Pick a stack, toggle the features you need, and get a ready-to-run project — with dependencies installed and boilerplate wired up.

<br/>

## Installation

```bash
git clone https://github.com/cmontecinos/forge.git
cd forge
go build -o forge ./cmd/forge

# Optional: add to PATH
sudo mv forge /usr/local/bin/
```

> **Requirements:** Go 1.21+, Git, Node.js 18+

<br/>

## Quick Start

```bash
# Interactive — prompts for stack and features
forge new my-app

# Non-interactive
forge new my-app --stack web --features auth,database,api

# Web-only (no Go backend)
forge new my-app --stack web-fullstack --features auth,database

# Start developing
cd my-app && npm run dev
```

<br/>

## Stacks

| | Stack | Frontend | Backend | Best for |
|---|-------|----------|---------|----------|
| :globe_with_meridians: | **`web`** | Next.js 14 (App Router) | Go + Echo | Fullstack web app with separate API |
| :iphone: | **`mobile`** | Expo / React Native | Go + Echo | Mobile app with separate API |
| :zap: | **`web-fullstack`** | Next.js 14 (App Router) | — | Web app talking to Supabase directly |

All stacks include **TypeScript**, **Tailwind CSS**, and **Supabase** out of the box.

<br/>

## Features

Optional modules you toggle during project creation:

| Feature | What you get | Stacks |
|---------|-------------|--------|
| **`auth`** | JWT auth flow — login, register, protected routes, auth context | `web` `mobile` `web-fullstack` |
| **`database`** | Supabase client, query builder, repository pattern, example model | `web` `mobile` `web-fullstack` |
| **`api`** | Echo router, middleware stack, CRUD handlers, error helpers | `web` `mobile` |

> :bulb: When using `web-fullstack`, the `api` feature is automatically excluded since there's no Go backend.

<br/>

## CLI Reference

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
forge version    # Show version info
forge --help     # Show help
```

<br/>

## Project Structure

### `web` / `mobile`

```
my-app/
├── frontend/ (or mobile/)
│   ├── src/
│   │   ├── app/              # Pages / screens
│   │   ├── components/       # UI components
│   │   └── lib/              # Utils & API client
│   └── package.json
├── backend/
│   ├── cmd/server/           # Entry point
│   ├── internal/             # Handlers, models, repos
│   └── go.mod
└── README.md
```

### `web-fullstack`

```
my-app/
├── src/
│   ├── app/                  # Next.js App Router
│   ├── components/           # UI components
│   └── lib/                  # Supabase client & utils
├── package.json
└── README.md
```

<br/>

## Environment Variables

<details>
<summary><strong>Backend</strong> (<code>.env</code>) — <code>web</code> and <code>mobile</code> stacks</summary>

```env
PORT=8080
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-anon-key
SUPABASE_SERVICE_KEY=your-service-key
JWT_SECRET=your-jwt-secret
```
</details>

<details>
<summary><strong>Frontend</strong> (<code>.env.local</code>)</summary>

```env
# web / mobile
NEXT_PUBLIC_API_URL=http://localhost:8080

# web-fullstack
NEXT_PUBLIC_SUPABASE_URL=https://your-project.supabase.co
NEXT_PUBLIC_SUPABASE_ANON_KEY=your-anon-key
```
</details>

<br/>

## Development

```bash
go build -o forge ./cmd/forge   # Build
go test ./...                    # Test
go vet ./...                     # Lint
```

**Project name rules:** starts with a letter, alphanumeric + hyphens + underscores, max 64 chars.

<br/>

## License

MIT
