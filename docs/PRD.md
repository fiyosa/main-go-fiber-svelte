# PRD — Go Fiber Portfolio Project

## 1. Project Overview

- **Name:** go-fiber-svelte
- **Module:** `go-fiber-svelte`
- **Stack:** Go 1.22+ (Fiber v2) + PostgreSQL (GORM) + Svelte 5 SPA (Vite + Tailwind CSS v4)
- **Project Layout:** [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- **Server:** Go Fiber (via `gofiber/fiber/v2`)
- **Deployment:** Docker (multi-stage: golang:alpine build + alpine runtime, port 8000)
- **Purpose:** Backend API with RBAC auth system, serving as a portfolio backend + SPA frontend

---

## 2. Project Structure

Project follows [golang-standards/project-layout](https://github.com/golang-standards/project-layout) conventions:

| Directory | Purpose |
|-----------|---------|
| `cmd/` | Main applications — one subdirectory per binary (`cmd/app`, `cmd/migrate`). Each has a small `main.go` that imports from `internal/`. |
| `internal/` | Private application code — **not importable by external projects** (enforced by Go compiler). All core logic lives here. |
| `web/` | Web application components — Vite + Svelte 5 SPA frontend. |
| `public/` | Build output and static assets served by Go Fiber. |
| `build/` | Packaging — Dockerfile in `build/package/`. |
| — | `.env.example` at project root. |
| `docs/` | Project documentation. |
| `scripts/` | Build/deploy scripts. |

> **Key rule:** No `src/` directory at the project root. Go's workspace (`$GOPATH`) has its own `src/`, but project-level `src/` is a Java anti-pattern.

```
.
.
├── go.mod                       # Module: go-fiber-svelte
├── go.sum
├── .gitignore
├── .air.toml                    # Hot-reload with go build (Windows .exe)
│
├── cmd/                         # Main applications
│   ├── app/
│   │   └── main.go              # Entry point: init config, DB, Fiber app, serve SPA
│   └── migrate/
│       └── main.go              # DB migration runner (CLI)
│
├── internal/                    # Private application code (enforced by Go compiler)
│   ├── bootstrap/
│   │   └── app.go               # Init middleware, register routes
│   │
│   ├── config/
│   │   └── config.go            # Load .env → Config struct
│   │
│   ├── controllers/             # Thin HTTP handlers
│   │   ├── auth_controller.go
│   │   ├── guest_controller.go
│   │   ├── doc_controller.go
│   │   └── policy_controller.go
│   │
│   ├── db/                      # Database layer (GORM)
│   │   ├── db.go                # Koneksi & init GORM (PostgreSQL)
│   │   ├── models/              # GORM model structs (7 tabel)
│   │   │   ├── user.go
│   │   │   ├── user_detail.go
│   │   │   ├── auth.go
│   │   │   ├── role.go
│   │   │   ├── permission.go
│   │   │   ├── role_has_permission.go
│   │   │   └── user_has_role.go
│   │   ├── migrations/          # AutoMigrate / SQL migration files
│   │   └── seed/
│   │       └── seed.go          # Seeder data awal
│   │
│   ├── helper/
│   │   └── response.go          # resSuccess, resSuccessData, resError, resPaginate, resCatch, resValidate
│   │
│   ├── lang/                    # i18n
│   │   ├── lang.go              # t._("key"), t._("key", args) helper
│   │   └── locales/
│   │       └── en.json          # English translations
│   │
│   ├── lib/                     # Reusable library packages
│   │   ├── hash.go              # bcrypt generate/verify + hashid encodeId/decodeId
│   │   ├── jwt.go               # JWT create/verify
│   │   ├── logger.go            # Logger (zerolog)
│   │   └── validator.go         # go-playground/validator wrapper
│   │
│   ├── middleware/
│   │   ├── auth_middleware.go    # Validasi JWT cookie → c.Locals("user_id")
│   │   └── hash_middleware.go   # Decode Hashids params ke numeric ID
│   │
│   ├── openapi/
│   │   └── openapi.go           # Generate OpenAPI spec
│   │
│   ├── provider/                # Core engine
│   │   ├── route_provider.go    # Route DSL helpers
│   │   ├── auth_provider.go     # RBAC: user_has_roles → role_has_permissions
│   │   └── app_provider.go      # Middleware registration
│   │
│   ├── repositories/            # Business logic (1 file per endpoint)
│   │   ├── auth/
│   │   │   ├── login_repository.go
│   │   │   ├── logout_repository.go
│   │   │   └── user_repository.go
│   │   └── policy/
│   │       ├── role_list_repository.go
│   │       ├── permission_list_repository.go
│   │       ├── permission_store_repository.go
│   │       └── permission_destroy_repository.go
│   │
│   ├── request/                 # Struct validasi (go-playground/validator)
│   │   ├── auth/
│   │   │   └── login_request.go
│   │   └── policy/
│   │       └── permission_store_request.go
│   │
│   ├── resources/               # Data transformers (encodeId tiap field id)
│   │   ├── auth/
│   │   │   └── user_resource.go
│   │   └── policy/
│   │       ├── role_list_resource.go
│   │       └── permission_resource.go
│   │
│   ├── routes/
│   │   └── api.go               # Route definitions & handler mapping
│   │
│   └── utils/
│       ├── date.go              # now(), formatByDate(), formatByStr()
│       └── uuid.go              # UUID v4 generate & validate
│
├── web/                         # Vite + Svelte 5 SPA
│   ├── index.html               # Entry HTML
│   ├── package.json
│   ├── vite.config.ts            # Vite + Tailwind v4 plugin + Svelte
│   ├── svelte.config.js
│   ├── tsconfig.json
│   ├── tsconfig.node.json
│   ├── .gitignore
│   ├── .prettierrc
│   ├── .vscode/
│   │   └── settings.json        # Prettier format-on-save for Svelte
│   │
│   ├── src/
│   │   ├── pages/
│   │   │   ├── routes.ts        # Route definitions (map with '*' catch-all)
│   │   │   └── guest/
│   │   │       ├── Home.svelte
│   │   │       ├── About.svelte
│   │   │       └── NotFound.svelte
│   │   │
│   │   ├── lib/                 # Svelte components & logic
│   │   │   ├── components/
│   │   │   │   └── ...svelte
│   │   │   ├── stores/
│   │   │   │   └── auth.ts
│   │   │   └── utils/
│   │   │       ├── axios.ts     # Axios instance + interceptor
│   │   │       └── ...
│   │   │
│   │   ├── api/                 # Frontend API hooks (TanStack Query)
│   │   │   ├── auth/
│   │   │   │   ├── index.ts
│   │   │   │   ├── login.ts
│   │   │   │   ├── logout.ts
│   │   │   │   └── user.ts
│   │   │   ├── guest/
│   │   │   │   ├── index.ts
│   │   │   │   └── ping.ts
│   │   │   └── policy/
│   │   │       ├── index.ts
│   │   │       ├── role_list.ts
│   │   │       ├── permission_list.ts
│   │   │       ├── permission_store.ts
│   │   │       └── permission_destroy.ts
│   │   │
│   │   ├── App.svelte           # Root: Router + routes
│   │   ├── main.ts              # Entry: setHashRoutingEnabled(false) + mount
│   │   ├── app.css              # Global styles @import "tailwindcss"
│   │   └── vite-env.d.ts
│   │
│   └── node_modules/            # pnpm dependencies (gitignored)
│
├── public/                      # Build output & static assets
│   ├── build/                   # Vite build output (served by Go Fiber)
│   │   ├── index.html
│   │   ├── main.js
│   │   └── main.css
│   ├── favicon.svg
│   ├── openapi.html             # Scalar API docs UI (local)
│   └── scalar-standalone.js     # Vendored Scalar JS (~3.6MB)
│
├── build/                       # Packaging & CI
│   └── package/
│       └── Dockerfile           # Multi-stage: build Go + build SPA
│
├── .env.example                 # Environment variable template
├── docs/
│   └── PRD.md                   # ← This file
│
└── scripts/
    └── deploy.sh                # Deployment script
```

---

## 3. Architecture & Patterns

### 3.1 Request Flow

```
HTTP Request
  → Fiber router (internal/routes/api.go)
    → Middleware chain (auth, hash)
      → Policy check (RBAC via authProvider)
        → Controller
          → Repository (business logic + GORM queries)
            → Resource (transform response)
              → Helper response (JSON)
```

### 3.2 Middleware

Middleware is registered per-route in `routes/api.go`, not globally. Public routes skip auth; protected routes use `AuthMiddleware` via route grouping.

**Existing middleware:**
| Name | File | Purpose |
|------|------|---------|
| `auth` | `auth_middleware.go` | Validates JWT cookie, sets `c.Locals("user_id")` |
| `hash` | `hash_middleware.go` | Decodes Hashids params to numeric IDs |

### 3.3 Policy/RBAC

Policies are checked via `authProvider.go`. Uses `user_has_roles → role_has_permissions → permissions` chain.

Applied per-route. Returns 403 if user lacks permission.

### 3.4 Controller → Repository Pattern

```
Controller (thin, delegates)
  → Repository (business logic, db queries)
    → Resource (data transformation)
```

Controllers are exported functions that delegate to repositories. Repositories are exported async functions.

### 3.5 Response Helpers

All from `internal/helper/response.go`:
| Function | Purpose |
|----------|---------|
| `resSuccess(msg, status)` | Success message |
| `resSuccessData(data, msg, status)` | Success with data |
| `resError(msg, errors, status)` | Error response |
| `resPaginate(data, meta, msg)` | Paginated response |
| `resCatch(error)` | Catch-block handler |
| `resValidate(error)` | Validation error |

**Standard response format:**
```json
{
  "message": "...",
  "data": { ... },
  "errors": { ... },
  "meta": { "total": 0, "page": 1, "limit": 10, "total_page": 0 }
}
```

---

## 4. Database Schema

All tables use `bigserial` PKs and GORM with relations.

### Tables:

| Table | Columns | Relations |
|-------|---------|-----------|
| `users` | id, email, username, password, created_at, updated_at, deleted_at | → auths, user_details, user_has_roles |
| `user_details` | id, user_id, first_name, last_name, created_at, updated_at | → users |
| `auths` | id, user_id, token, revoke, ip, user_agent, created_at, updated_at | → users |
| `roles` | id, name, notes, created_at, updated_at, deleted_at | → user_has_roles, role_has_permissions |
| `permissions` | id, name, notes, created_at, updated_at, deleted_at | → role_has_permissions |
| `role_has_permissions` | role_id, permission_id (composite PK) | → roles, permissions |
| `user_has_roles` | user_id, role_id (composite PK) | → users, roles |

### Key relationships:
- User M:N Role via `user_has_roles`
- Role M:N Permission via `role_has_permissions`
- All junction tables use composite primary keys
- `deleted_at` used for soft-delete on users, roles, permissions

### GORM client:
```go
// internal/db/db.go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info), // local only
})
```

---

## 5. API Endpoints

| Method | Path | Auth | Policy | Controller | Description |
|--------|------|------|--------|------------|-------------|
| GET | `/api/openapi.json` | - | - | `docController.openapi` | OpenAPI 3.0 spec (dynamic) |
| GET | `/api/docs` | - | - | `docController.docs` | Scalar API docs UI |
| POST | `/api/auth/login` | - | - | `authController.login` | Login, HTTP-only cookie |
| DELETE | `/api/auth/logout` | auth | - | `authController.logout` | Revoke token + clear cookie |
| GET | `/api/auth/user` | auth | - | `authController.user` | Current user info |
| GET | `/api/policy/role` | auth | - | `policyController.roleList` | List roles with permissions |
| GET | `/api/policy/permission` | auth | - | `policyController.permissionList` | List permissions |
| POST | `/api/policy/permission` | auth | - | `policyController.permissionStore` | Create permission |
| DELETE | `/api/policy/permission/:id` | auth | - | `policyController.permissionDestroy` | Delete permission (soft) |
| GET | `/api/guest/ping` | - | - | `guestController.ping` | Health check |
| * | `/api/*` | - | - | 404 JSON | Fallback for invalid API paths |

---

## 6. Libraries & Utilities

### Go Libraries (`internal/lib/`):

| File | Exports | Purpose |
|------|---------|---------|
| `hash.go` | `Generate`, `Verify`, `EncodeId`, `DecodeId` | bcrypt hashing + Hashids encoding |
| `jwt.go` | `Create`, `Verify` | JWT sign/verify with APP_SECRET |
| `logger.go` | `Info`, `Error`, `Warn`, `Debug` | Zerolog logging |
| `validator.go` | `Validate`, `ValidateRequest` | go-playground/validator wrapper |

### Go Utilities (`internal/utils/`):

| File | Exports | Purpose |
|------|---------|---------|
| `date.go` | `Now`, `FormatByDate`, `FormatByStr` | Date formatting |
| `uuid.go` | `Create`, `Verify` | UUID v4 generation & validation |

### Frontend Dependencies (SPA):

| Package | Version | Purpose |
|---------|---------|---------|
| `svelte` | ^5.55.5 | UI framework (runes-based) |
| `@keenmate/svelte-spa-router` | ^5.1.0 | SPA router (history mode with `use:link`) |
| `@tailwindcss/vite` | ^4.3.0 | Tailwind CSS v4 Vite plugin |
| `tailwindcss` | ^4.3.0 | Utility-first CSS |
| `vite` | ^8.0.12 | Build tool |
| `@sveltejs/vite-plugin-svelte` | ^7.1.2 | Svelte Vite integration |
| `prettier` | ^3.8.3 | Code formatter |
| `prettier-plugin-svelte` | ^4.1.0 | Svelte Prettier plugin |
| `prettier-plugin-tailwindcss` | ^0.8.0 | Tailwind class sorting |

---

## 7. Configuration & Environment

### Environment Variables (.env):

| Variable | Example | Scope |
|----------|---------|-------|
| `APP_ENV` | `local` | Public |
| `APP_LOCALE` | `en` | Public |
| `APP_SECRET` | `secret` | Secret |
| `APP_JWT_DURATION` | `1d` | Public |
| `DB_URL` | `postgresql://...` | Secret |
| `API_URL` | `http://localhost:8000` | Public |
| `PORT` | `8000` | Public |

Config loaded via `internal/config/config.go` (using `github.com/joho/godotenv` + `os.Getenv` or `caarlos0/env`).

---

## 8. Internationalization

Custom i18n system at `internal/lang/`:
- `t._("key")` → returns translated string
- `t._("key", map[string]interface{}{"arg": "value"})` → replaces `:arg` in string
- Currently only `en` locale exists
- `APP_LOCALE` controls active language

---

## 9. Code Conventions

### Go:

- **Controllers:** Exported functions that delegate to repositories. For multiple methods on same resource, use a struct grouping or separate functions. Barrel pattern via package-level exports.
- **Repositories:** Exported functions; **1 file per API endpoint** with specific name (`RoleListRepository`, `PermissionStoreRepository`). Barrel via `index.go` per subfolder.
- **Resources:** Exported functions `Single()` dan `Collection()`, bukan struct method. Setiap field `id` wajib di-encode dengan `hash.EncodeId()`. Barrel via `index.go` per subfolder.
- **Request schemas:** Struct validation in `src/request/`, barrel via `index.go` per subfolder.
- **Middleware:** Exported function + registered via `appProvider`.
- **Handlers:** Return `c.JSON()` responses via helper functions.

### Frontend (Svelte/TypeScript):

- **Router:** `@keenmate/svelte-spa-router` with **history mode** (`setHashRoutingEnabled(false)`). Hash redirect fallback converts `/#/xyz` → `/xyz`.
- **Route definitions:** Map `{ '/': Home, '/about': About, '*': NotFound }` or `defineRoutes()` for type-safe named routes. Catch-all `'*'` matches unmatched paths.
- **Navigation links:** `<a href="/path" use:link>` for SPA navigation. Programmatic: `push(path)` / `nav.about.push()`.
- **Route params:** `{ path: '/user/:id' }` → component receives `let { routeParams = {} } = $props()` — access via `routeParams.id`.
- **Query string helpers:** `query<T>()` for reactive query params, `updateQuerystring()` for modifications.
- **Navigation guards:** `registerBeforeLeave()` for dirty-form protection.
- **Permissions (frontend):** Built-in RBAC with `createProtectedRoute()`, `hasPermission()`.
- **API hooks:** `web/src/api/` berisi wrapper TanStack Query per endpoint. Naming: `{method}{Name}` (getUser, postLogin, etc). Barrel per subfolder.
- **TypeScript:** Strict mode.
- **Formatting:** Prettier with `prettier-plugin-svelte` + `prettier-plugin-tailwindcss`.
- **CSS:** Tailwind CSS v4 via `@tailwindcss/vite` plugin — import `@import "tailwindcss"` di `app.css`, class-based styling, otomatis sort via Prettier.
- **Imports:** Use relative paths within `web/src/`.

### Tooling & IDE:

- **VS Code:** `editor.formatOnSave: true`, Svelte formatter set ke Prettier via `.vscode/settings.json`.
- **Go hot-reload:** Air (`air`) configured in `.air.toml` (Windows-compatible: `.exe` extension). Watches `cmd/` + `internal/`, builds `./cmd/app`.

---

## 10. NPM Scripts (SPA)

| Script | Command |
|--------|---------|
| `dev` | `vite --host` (port 5173 / configured) |
| `build` | `vite build` → output `../../public/build/` |
| `preview` | `vite preview --host` |
| `check` | `svelte-check --tsconfig ./tsconfig.app.json && tsc -p tsconfig.node.json` |
| `format` | `prettier --write 'src/**/*.{svelte,ts,css}'` |

---

## 11. Go Scripts

| Script | Command | Port |
|--------|---------|------|
| `dev` | `go run ./cmd/app` / `air` (hot-reload) | 8000 |
| `build` | `go build -o app ./cmd/app` | - |
| `run` | `./app` | 8000 |
| `migrate` | `go run ./cmd/migrate` | - |
| `lint` | `golangci-lint run ./...` | - |
| `test` | `go test ./...` | - |

Hot-reload via `air` (`.air.toml`): watches `cmd/` + `internal/`, rebuilds to `./tmp/main.exe` from `./cmd/app`.

### CLI Commands

- **`go run ./cmd/app`** — Start API server
- **`go run ./cmd/migrate`** — Run database migrations (AutoMigrate)

---

## 12. Production

### Build & Deploy

```bash
# SPA build
cd web && pnpm install && pnpm build

# Go build
cd ..
go build -o app ./cmd/app

# Run
./app
```

### main.go production logic (`cmd/app/main.go`)

```go
// Static assets (CSS/JS from Vite build)
app.Static("/", "public")

// SPA catch-all — serve index.html for all non-API paths
app.Get("/api/*", func(c *fiber.Ctx) error {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
        "message": "API endpoint not found",
    })
})
app.Get("/*", func(c *fiber.Ctx) error {
    return c.SendFile("public/index.html")
})
```

### Docker

Multi-stage build:
```
Stage 1 (node:20-alpine) → pnpm install && pnpm build (from web/) → public/build/
Stage 2 (golang:1.22-alpine) → go build ./cmd/app → app
Stage 3 (alpine) → copy app + public/build/ → ./app (port 8000)
```

### SPA Development Mode

When `APP_ENV=local`, SPA runs separately on `:3000` (Vite dev server). Production uses built files served by Go Fiber from `public/`. To start SPA dev: `cd web && pnpm dev`.

### SPA Router

| Package | `@keenmate/svelte-spa-router` (fork with Svelte 5 runes) |
|---------|----------|
| **Mode** | History mode via `setHashRoutingEnabled(false)` + `setBasePath('/')` in `main.ts` |
| **Hash redirect** | `/#/xyz` → `history.replaceState` to `/xyz` on init |
| **Link navigation** | `<a href="/path" use:link>` or programmatic `push(path)` |
| **Params** | `{ path: '/user/:id' }` → `routeParams.id` in component |
| **Catch-all** | `'*': NotFound` in route map for 404 pages |

### API Documentation (OpenAPI / Scalar)

- OpenAPI 3.0 spec generated dynamically at `/api/openapi.json`
- Per-controller `*OpenAPIPaths()` functions merged in `internal/openapi/openapi.go`
- Scalar UI served at `/api/docs` using vendored `public/scalar-standalone.js` (offline-capable)

### Go Middleware

**Auth middleware** (`internal/middleware/auth_middleware.go`): validates JWT cookie, sets `c.Locals("user_id")`. Applied **per-route** (not globally) via `internal/routes/api.go`:

```go
api := app.Group("/api")
authMw := AuthMiddleware{...}

// Public routes (no auth)
api.Post("/auth/login", authController.Login)
api.Get("/guest/ping", guestController.Ping)

// Protected routes (with auth)
auth := api.Group("/auth", authMw.Handle)
auth.Get("/user", authController.User)
```

Public API routes: `/api/openapi.json`, `/api/docs`, `/api/guest/ping`, `/api/auth/login`. All others return 401 if unauthenticated.
