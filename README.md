# Disc golf watch scorekeeper

A learning project for Go, PostgreSQL, Astro, smartwatch development, and DevOps.
This repository contains a runnable skeleton, not a finished scorekeeper.

## Structure

```text
backend/
  cmd/api/               API entry point and dependency wiring
  internal/config/       Environment configuration
  internal/httpapi/      Router and future handlers
  internal/rounds/       Future models and application rules
  internal/postgres/     Future persistence code
  migrations/            Instructions for adding schema migrations
  examples/telemetry/    Original prototype, preserved separately
frontend/src/
  pages/                 Existing portfolio and dashboard placeholder
  layouts/               Shared page structure
  components/            UI components
  lib/api.ts             Starting point for browser API calls
watch/                   Watch client planning placeholder
docs/
  architecture.md        Package responsibilities and open decisions
  api.yaml               Empty OpenAPI contract to develop yourself
```

Start with [the architecture guide](docs/architecture.md). TODO comments mark
places for your implementation. There are no predefined round fields, scoring
rules, storage interfaces, schemas, or application endpoints.

## Run locally

Use the Go version declared in `backend/go.mod`, Node.js matching
`frontend/package.json`, and Docker Compose if you need PostgreSQL.

Start the API in one terminal:

```sh
cd backend
go run ./cmd/api
```

It listens on port 8080. All paths return 404 until you register routes in
`internal/httpapi/router.go`. PostgreSQL is not required for this skeleton.
Stop it with Ctrl+C. The old `go run main.go` command has been replaced.

To override the address:

```sh
HTTP_ADDR=127.0.0.1:8081 go run ./cmd/api
```

The root `.env.example` documents configuration. The Go app does not load `.env`
files automatically; pass environment variables through your shell or runner.

Start the frontend in another terminal:

```sh
cd frontend
npm ci
npm run dev -- --background
```

Open the local URL reported by Astro and visit `/dashboard/` for the placeholder.
Manage the background server with `npm run astro -- dev status`,
`npm run astro -- dev logs`, and `npm run astro -- dev stop`.

When you are ready to work with PostgreSQL, run from the repository root:

```sh
docker compose up -d db
```

The existing Compose settings and database volume are unchanged. No migrations
are executed and the new API does not read or write the database yet.
The original telemetry example has [separate run instructions](backend/examples/telemetry/README.md).

## Checks

From `backend/`:

```sh
go test ./...
go vet ./...
```

There are no tests yet. Add meaningful tests alongside your implementation.
Use `gofmt -w` on Go files you edit.

From `frontend/`:

```sh
npm run build
```

CI runs Go formatting checks, tests, vet, and the Astro build on pull requests
and pushes to `main`. An Astro build is not a full TypeScript type check.
The existing Pages deployment workflow remains separate; CI does not gate it.

## Suggested first implementation

1. Define the data needed for one round in `internal/rounds/model.go`.
2. Describe one operation in `docs/api.yaml`.
3. Implement its service and tests, then define the storage interface it needs.
4. Choose a migration tool, add a schema, and implement persistence.
5. Wire the handler and dependencies in `cmd/api/main.go`.
6. Implement the browser API call and dashboard component.

Choose watch platform, authentication, synchronization rules, and deployment
details when you are ready to work on those parts.
