# Forge

## What This Is

CLI interactivo en Go para crear proyectos con stacks personalizados. Elimina la fricción de configurar boilerplate cada vez que inicias un proyecto nuevo, generando estructuras completas con tus configuraciones exactas.

## Core Value

**Mis configuraciones exactas, no genéricas** — genera proyectos con MI stack preferido, listos para agregar funcionalidad, sin configuración manual.

## Requirements

### Validated

(None yet — ship to validate)

### Active

- [ ] CLI interactivo en Go con flujo: nombre → tipo → features
- [ ] Stack Web: Next.js (App Router + Tailwind + TS) + Go (Echo + project-layout) + Supabase
- [ ] Stack Mobile: Expo (React Native) + Go (Echo + project-layout) + Supabase
- [ ] Estructura monorepo: `/frontend` o `/mobile` + `/backend`
- [ ] Arquitectura: Frontend → Go → Supabase (nunca conexión directa)
- [ ] Feature Auth: Login/registro via backend, validación de tokens en Go
- [ ] Feature Database: Conexión Go-Supabase, modelos base
- [ ] Feature API: Router Echo, middlewares, handlers de ejemplo
- [ ] Templates almacenados en Git repos (versionables, editables)
- [ ] Post-create: instalación automática de dependencias
- [ ] Arquitectura extensible: stacks y features como módulos independientes
- [ ] Detección dinámica de stacks y features disponibles (no hardcodeado)

### Out of Scope

- Otros stacks/frameworks — v1 solo soporta los dos stacks definidos
- Interfaz web — es CLI personal, no necesita UI
- Distribución pública — herramienta personal, no para publicar
- git init automático — el usuario lo hace manualmente si quiere
- Abrir en editor — el usuario decide qué editor usar

## Context

**Motivación:** El usuario construye software para resolver sus propios problemas. Cada nuevo proyecto requiere crear boilerplate desde cero, lo cual es tiempo perdido y propenso a inconsistencias.

**Stacks definidos:**
- Web: Next.js + Go + Supabase
- Mobile: React Native (Expo) + Go + Supabase

**Patrón arquitectónico:** El frontend nunca habla directamente con Supabase. Todo pasa por el backend Go, que actúa como única puerta a la base de datos y autenticación.

**Estructura Go:**
- Framework: Echo
- Layout: golang-standards/project-layout (cmd/, internal/, pkg/, api/)

**Estructura Next.js:**
- App Router (Next.js 13+)
- Tailwind CSS
- TypeScript

**Estructura React Native:**
- Expo (managed workflow)

**Extensibilidad:** El diseño debe permitir agregar nuevos stacks y features fácilmente. Estructura modular donde cada stack y feature es un módulo independiente.

## Constraints

- **Tech stack CLI**: Go — consistente con los backends que genera
- **Templates**: Git repos — permite versionado y edición fácil
- **Arquitectura**: Frontend → Go → Supabase — nunca conexión directa desde frontend

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| CLI en Go | Consistente con stack backend, binario único | — Pending |
| Templates en Git repos | Versionables, editables, potencial para compartir futuro | — Pending |
| Monorepo structure | Simplifica desarrollo, un solo repo por proyecto | — Pending |
| Backend como proxy a Supabase | Seguridad, control centralizado, flexibilidad | — Pending |
| Echo framework | Familiar para el usuario, buen balance simplicidad/features | — Pending |
| Expo para React Native | Simplifica builds y desarrollo mobile | — Pending |
| Stacks/features como módulos | Extensibilidad sin modificar core | — Pending |

---
*Last updated: 2025-01-14 after initialization*
