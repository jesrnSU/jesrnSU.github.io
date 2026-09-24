# Architecture skeleton

The repository keeps the existing Go/Gin/pgx backend, PostgreSQL development
container, and static Astro frontend. The scaffold introduces package boundaries
without implementing a disc golf data model or application behavior.

## Responsibilities

| Location | Responsibility |
| --- | --- |
| `backend/cmd/api` | Startup, dependency wiring, and eventually shutdown |
| `backend/internal/config` | Environment configuration |
| `backend/internal/httpapi` | HTTP routing, decoding, and responses |
| `backend/internal/rounds` | Your data model, use cases, and scoring rules |
| `backend/internal/postgres` | SQL and transaction handling |
| `backend/migrations` | Versioned database schema changes |
| `frontend/src/pages` | Page routes |
| `frontend/src/layouts` | Shared page structure |
| `frontend/src/components` | UI components as features are added |
| `frontend/src/lib/api.ts` | Future browser-side API calls |
| `watch` | Future watch client |
| `docs/api.yaml` | API contract as you define it |

The intended call direction is HTTP handler → service → repository. Define
storage interfaces in the service package when a concrete use case needs them;
implement them in `postgres`. Wire the implementations in `cmd/api`.

## Decisions left open

- Round fields, player relationships, identifiers, and score validation.
- Endpoint paths, request/response shapes, and error conventions.
- Database schema, migration tooling, and transaction boundaries.
- Offline storage, retry semantics, and handling conflicting edits.
- Authentication, ownership checks, and browser CORS configuration.
- Statistics definitions, charting, and dashboard design.
- Watch platform and backend hosting/container deployment.

No placeholder operation reports success or returns invented data. The new API
has no registered endpoints and does not connect to PostgreSQL. The old telemetry
prototype is isolated under `backend/examples/telemetry`.

## Frontend deployment

The existing workflow still publishes Astro to GitHub Pages. The new dashboard
uses Astro's [file-based routing](https://docs.astro.build/en/guides/routing/)
and [components](https://docs.astro.build/en/basics/astro-components/).
Live data will need browser-side requests to the Go API; static page frontmatter
runs during the build. See [Astro client scripts](https://docs.astro.build/en/guides/client-side-scripts/).

The API and database have no deployment configuration yet. Add a backend
Dockerfile and deployment workflow once you choose how to run and configure them.
The existing portfolio project brief remains aspirational; it is not a list of
features implemented by this scaffold.
