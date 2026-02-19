# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Forge is a Go CLI tool for scaffolding new projects. It clones Git template repositories, processes files as Go `text/template` templates (substituting `{{.ProjectName}}`, `{{.StackType}}`, `{{.Features}}`), and optionally runs post-create dependency installation.

Two stacks: **Web** (Next.js 14 + Go/Echo + Supabase) and **Mobile** (Expo/React Native + Go/Echo + Supabase). Three optional features: auth, database, api. Template content lives in external repos — this repo is only the generator CLI.

## Build & Run

```bash
# Build
go build -o forge ./cmd/forge

# Build with version metadata
go build -ldflags "-X github.com/bigbytes/forge/internal/config.Version=1.0.0 \
  -X github.com/bigbytes/forge/internal/config.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
  -X github.com/bigbytes/forge/internal/config.GitCommit=$(git rev-parse --short HEAD)" \
  -o forge ./cmd/forge

# Run
./forge new my-app
./forge new my-app --stack web --features auth,api
./forge version

# Test
go test ./...

# Vet
go vet ./...
```

## Architecture

Entry point: `cmd/forge/main.go` → calls `cli.Execute()`.

**Key packages under `internal/`:**

- **cli/** — Cobra commands (`root.go`, `new.go`, `version.go`), interactive prompts (`prompts.go`), colored output helpers (`output.go`), error types (`errors.go`), post-create hooks (`postcreate.go`)
- **stacks/** — `Stack` interface + thread-safe registry. Implementations self-register via `init()` (web.go, mobile.go)
- **features/** — `Feature` interface + thread-safe registry. Same self-registration pattern. Each feature declares compatible stacks via `CompatibleStacks()`
- **templates/** — Git clone (shallow, depth=1), template variable substitution across all text files, `.git` cleanup. `TemplateError` wraps errors with operation context
- **config/** — Version/BuildDate/GitCommit vars set via ldflags at build time

## Key Patterns

- **Self-registering modules**: Stacks and features register themselves in `init()` functions. To add a new stack/feature: create a file, implement the interface, call `Register()` in `init()`. No central switch to update.
- **Thread-safe registries**: Both registries use `sync.RWMutex`.
- **Template processing**: Binary files (containing null bytes) and files without `{{`/`}}` markers are skipped.
- **Features are metadata-only in this repo**: Feature logic lives in the template repos. The CLI passes feature names in `TemplateData` for conditional template rendering.
- **Table-driven tests**: Standard Go test pattern used throughout.
