# Roadmap: Forge

## Overview

Forge es un CLI en Go para crear proyectos con stacks personalizados. El desarrollo sigue una progresión desde la infraestructura del CLI, pasando por los sistemas modulares de stacks y features, hasta los templates concretos y su integración final. El objetivo es tener una herramienta funcional que genere proyectos Next.js/Expo + Go + Supabase listos para desarrollar.

## Domain Expertise

None

## Phases

**Phase Numbering:**
- Integer phases (1, 2, 3): Planned milestone work
- Decimal phases (2.1, 2.2): Urgent insertions (marked with INSERTED)

- [x] **Phase 1: CLI Foundation** - Estructura base del CLI, parsing de argumentos, flujo interactivo
- [ ] **Phase 2: Template Engine** - Sistema para clonar y procesar templates desde Git repos
- [ ] **Phase 3: Stack Registry** - Sistema modular de registro de stacks con detección dinámica
- [ ] **Phase 4: Feature Registry** - Sistema modular de features componibles entre stacks
- [ ] **Phase 5: Web Stack Base** - Template base Next.js + Go + Supabase (monorepo)
- [ ] **Phase 6: Mobile Stack Base** - Template base Expo + Go + Supabase (monorepo)
- [ ] **Phase 7: Auth Feature** - Módulo de autenticación via backend Go
- [ ] **Phase 8: Database Feature** - Módulo de conexión Go-Supabase y modelos base
- [ ] **Phase 9: API Feature** - Módulo de estructura API con Echo
- [ ] **Phase 10: Integration & Polish** - Flujo completo, instalación de deps, testing E2E

## Phase Details

### Phase 1: CLI Foundation
**Goal**: CLI funcional en Go con flujo interactivo: nombre → tipo → features
**Depends on**: Nothing (first phase)
**Research**: Unlikely (established Go CLI patterns)
**Plans**: TBD

Deliverables:
- Comando `forge new <project-name>`
- Prompts interactivos para selección de stack y features
- Estructura base del proyecto Go (cmd/, internal/)
- Help y version commands

### Phase 2: Template Engine
**Goal**: Sistema para clonar repos Git y procesar templates con variables
**Depends on**: Phase 1
**Research**: Likely (git operations in Go)
**Research topics**: go-git library, template processing, variable substitution
**Plans**: TBD

Deliverables:
- Clonar repos de templates desde URLs configurables
- Procesar archivos con variables ({{.ProjectName}}, etc.)
- Copiar estructura al directorio destino
- Manejo de errores y cleanup

### Phase 3: Stack Registry
**Goal**: Sistema modular para registrar y descubrir stacks disponibles
**Depends on**: Phase 2
**Research**: Unlikely (internal patterns)
**Plans**: TBD

Deliverables:
- Interfaz Stack con métodos estándar
- Registro dinámico de stacks
- Detección de stacks disponibles en runtime
- Configuración por stack (repo URL, variables requeridas)

### Phase 4: Feature Registry
**Goal**: Sistema modular para features componibles que aplican a múltiples stacks
**Depends on**: Phase 3
**Research**: Unlikely (internal patterns)
**Plans**: TBD

Deliverables:
- Interfaz Feature con métodos estándar
- Registro dinámico de features
- Compatibilidad feature-stack (qué features aplican a qué stacks)
- Sistema de aplicación de features sobre base stack

### Phase 5: Web Stack Base
**Goal**: Template funcional Next.js + Go + Supabase en estructura monorepo
**Depends on**: Phase 4
**Research**: Likely (current Next.js patterns)
**Research topics**: Next.js 14 App Router setup, golang-standards/project-layout, monorepo structure
**Plans**: TBD

Deliverables:
- Repo template: `/frontend` (Next.js App Router + Tailwind + TS)
- Repo template: `/backend` (Go Echo + project-layout)
- Configuración base de ambos
- Scripts de desarrollo (dev, build)

### Phase 6: Mobile Stack Base
**Goal**: Template funcional Expo + Go + Supabase en estructura monorepo
**Depends on**: Phase 5 (reutiliza backend)
**Research**: Likely (current Expo patterns)
**Research topics**: Expo SDK latest, recommended project structure, TypeScript setup
**Plans**: TBD

Deliverables:
- Repo template: `/mobile` (Expo managed workflow + TS)
- Reutilización de `/backend` del web stack
- Configuración Expo actual
- Scripts de desarrollo

### Phase 7: Auth Feature
**Goal**: Módulo de autenticación completo (Frontend → Go → Supabase)
**Depends on**: Phase 5, Phase 6
**Research**: Likely (Supabase Auth integration)
**Research topics**: Supabase Auth Go SDK, JWT validation in Echo, session handling patterns
**Plans**: TBD

Deliverables:
- Backend: endpoints auth (login, register, refresh, logout)
- Backend: middleware de validación JWT
- Frontend Next.js: hooks y componentes auth
- Frontend Expo: hooks y componentes auth
- Protección de rutas en ambos frontends

### Phase 8: Database Feature
**Goal**: Módulo de conexión Go-Supabase con modelos base
**Depends on**: Phase 7
**Research**: Likely (supabase-go client)
**Research topics**: supabase-go library, connection pooling, model patterns
**Plans**: TBD

Deliverables:
- Backend: cliente Supabase configurado
- Backend: modelos base de ejemplo
- Backend: repository pattern para queries
- Configuración de variables de entorno

### Phase 9: API Feature
**Goal**: Estructura API completa con Echo (router, middlewares, handlers)
**Depends on**: Phase 8
**Research**: Unlikely (established Echo patterns)
**Plans**: TBD

Deliverables:
- Backend: estructura de router con grupos
- Backend: middlewares estándar (CORS, logging, recovery)
- Backend: handlers de ejemplo (CRUD)
- Backend: error handling consistente

### Phase 10: Integration & Polish
**Goal**: Flujo completo funcional, instalación de deps, testing E2E
**Depends on**: Phase 9
**Research**: Unlikely (internal testing)
**Plans**: TBD

Deliverables:
- Flujo completo: `forge new` → proyecto listo
- Post-create: `npm install` / `go mod tidy` automático
- Testing E2E del CLI
- Documentación de uso (README)

## Progress

**Execution Order:**
Phases execute in numeric order: 1 → 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9 → 10

| Phase | Plans Complete | Status | Completed |
|-------|----------------|--------|-----------|
| 1. CLI Foundation | 5/5 | Complete | 2026-01-14 |
| 2. Template Engine | 0/TBD | Not started | - |
| 3. Stack Registry | 0/TBD | Not started | - |
| 4. Feature Registry | 0/TBD | Not started | - |
| 5. Web Stack Base | 0/TBD | Not started | - |
| 6. Mobile Stack Base | 0/TBD | Not started | - |
| 7. Auth Feature | 0/TBD | Not started | - |
| 8. Database Feature | 0/TBD | Not started | - |
| 9. API Feature | 0/TBD | Not started | - |
| 10. Integration & Polish | 0/TBD | Not started | - |
