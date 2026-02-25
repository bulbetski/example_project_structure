# Repository Guidelines

## Memory Bank
- Read `docs/agent/GUIDELINES.md` before any implementation work.
- Use `docs/` as the source of truth for structure and architecture.
- If documentation conflicts with code, report it and stop.
- If something is unclear, mark it `UNKNOWN`.

## Structure
- Keep service entrypoints in `cmd/app/main.go`.
- Keep subcommand implementations in `cmd/app/internal/*.go` and register them on `RootCmd` (cobra).
- Keep core application code in `internal/` with layer separation: transport (`httptransport`, `rpctransport`, `kafkatransport`), `service`, `repository`, `domain`.
- Keep dependency wiring in `internal/cli/deps` (define providers in `providers.go`, register them in `container.go`, and update `services.go`, `repositories.go`, `servers.go`, `kafka-consumers.go`, `cronjobs.go`, `migrate.go` as needed).
- For dependency initialization, use `samber.Do`.
- Keep configuration constants in `internal/config`.
- Keep internal-only helpers in `internal/pkg`; keep reusable helpers in `pkg`.
- Keep specs for codegen in `docs/protobuf` and `docs/openapi`.
- Keep runtime tooling and SQL migrations in `build/` (`build/local` for runtime configs, `build/app/clickhouse_migrations/*` for SQL).

## Layering And Dependencies
- Follow the existing dependency direction: transport -> service -> repository.
- Keep domain types in `internal/domain/*` and import them from `service` and `repository` where needed.
- Keep transport handlers dependent on `internal/service` (via `service.SvcLayer`) and generated transport interfaces.

## Subcommands
- Add new subcommands under `cmd/app/internal` and register them with `RootCmd.AddCommand`.
- Use the DI container only inside subcommand `main` execution; do not rely on the container for lazy runtime resolution.

## Migrations
- Add ClickHouse migrations under `build/app/clickhouse_migrations/{common,dev,uat,prod}`.
- Use `golang-migrate` naming: `<version>_<name>.up.sql` and `<version>_<name>.down.sql`.
- Keep version prefixes unique across `common` and the environment folder used at runtime.

## Build And Test Commands
- `make install` to install CLI helpers.
- `make up` / `make down` / `make status` for local dependencies.
- `make run` to start the service with `build/local/.env`.
- `make gen-mocks` to regenerate mocks.
- `make codegen-update` to refresh generated code.
- `make lint` for Go lint.
- `make test` and `make test-cov` for tests and coverage.
- `make load-test` for load tests.

## Coding And Naming
- Use `gofmt`/`goimports` with tab-based indentation.
- Keep package names short and lower_snake_case.
- Use PascalCase for exported identifiers and add doc comments.
- Group files by responsibility (e.g., `internal/service`, `internal/repository`).
- Wrap errors with `fmt.Errorf("obj.Method: %w", err)` or `fmt.Errorf("funcName: %w", err)`.

## Testing
- Place tests alongside sources (`file.go` -> `file_test.go`).
- Name tests `TestComponent_Scenario`.
- Use `t.Context()` in tests (not `context.Background()`).
- Follow ZOMBIES for test design: Zero, One, Many, Boundary, Interface, Exceptional, Simple.
- Use `require` from `github.com/stretchr/testify/require`.
- Use `go.uber.org/mock/gomock` for interaction-based tests.
- For interfaces used by a unit, add `//go:generate go run go.uber.org/mock/mockgen` near the interface, generate mocks into a `mocks/` subfolder, and run `go generate ./...` (or the package).
