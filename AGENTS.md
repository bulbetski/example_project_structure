# Repository Guidelines

## Structure

### Project Layout And Responsibilities

- `cmd/orders/main.go` is the service entrypoint.
- Subcommands live in `cmd/orders/internal` and are registered on `RootCmd` (cobra).
- `internal/cli/deps` owns dependency providers and DI container wiring used by subcommands:
    - define providers in `providers.go`
    - register them in `container.go`
    - update wiring surfaces as needed: `services.go`, `repositories.go`, `servers.go`, `kafka-consumers.go`,
      `cronjobs.go`, `migrate.go`
- `internal/config` holds configuration constants.
- `internal/domain` contains domain types shared by `internal/service` and `internal/repository`.
- `internal/service` implements service logic and defines interfaces for storage dependencies.
- `internal/repository` implements storage interfaces used by `internal/service`.
- `internal/httptransport` and `internal/rpctransport` provide HTTP and gRPC transport handlers.
- `internal/kafkatransport` provides Kafka consumer handlers.
- `internal/pkg` contains internal-only helpers; `pkg` contains reusable helpers.
- `build/local` keeps runtime configs and docker compose for local dependencies.
- `docs/openapi` and `docs/protobuf` store specs for code generation.
- `docs/agent` stores agent operating rules and structure guidance (including `AGENTS.md`, `GUIDELINES.md`,
  `STRUCTURE.md`).
- `scripts/system` and `scripts/usr` provide Make targets (included by the root `Makefile`).
- `tests/load` hosts load tests.

### Dependency Flow (Expected)

- `cmd/main`
- `cmd/internal` (cobra)
- `internal/cli/deps` (providers + container)
- transport (`httptransport` / `rpctransport` / `kafkatransport`)
- `internal/service`
- `internal/repository`

Additional notes:

- `internal/domain` types are imported by `internal/service` and `internal/repository`.
- `internal/config` is imported by wiring and runtime setup.

### DI Conventions

- For dependency initialization, use `samber.Do`.
- Keep the DI container usage scoped to subcommand execution (`main` path); do not rely on the container for lazy
  runtime resolution inside business code.

## Layering And Dependencies

- Follow the existing dependency direction: transport -> service -> repository.
- Keep domain types in `internal/domain/*` and import them from `service` and `repository` where needed.
- Keep transport handlers dependent on `internal/service` (via `service.SvcLayer`) and generated transport interfaces.

## Subcommands

- Add new subcommands under `cmd/orders/internal` and register them with `RootCmd.AddCommand`.
- Use the DI container only inside subcommand `main` execution; do not rely on the container for lazy runtime
  resolution.

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
- For interfaces used by a unit, add `//go:generate go run go.uber.org/mock/mockgen` near the interface, generate mocks
  into a `mocks/` subfolder, and run `go generate ./...` (or the package).

# Agent Operating Rules

## General Principles

- Do not write production code without a plan.
- If documentation conflicts with code, report it before proceeding.
- If something is unclear, mark it `UNKNOWN` instead of guessing.

## Workflow (Mandatory)

### 1) Context

Before implementing anything:

- Identify affected modules and boundaries.

### 2) Plan

Before writing code:

- Produce a structured implementation plan.
- List affected files.
- List new files to create.
- Identify dependency impacts.
- Identify required tests.
- Wait for confirmation (unless explicitly told to proceed automatically).

### 3) Implementation

- Follow architectural layering rules strictly.
- Do not violate dependency direction.
- Do not introduce cross-layer shortcuts.
- Keep changes minimal and scoped.

### 4) Testing

After implementation:

- Run tests.
- Add missing tests if required.
- Ensure lint/build passes.
- If tests fail, fix before proceeding.

### 5) Verification

Before declaring task complete:

- Check each item from the plan.
- Confirm it is implemented.
- Report file paths for each change.
- Report test results.

### 6) Git Discipline

- Commit after plan approval.
- Commit after implementation.
- Commit after tests pass.
- Commit after every task you do.
- Keep commits atomic and descriptive.

## Forbidden

- Writing code without plan.
- Skipping tests.
- Ignoring `AGENTS.md`.
- Modifying unrelated modules.
- Making assumptions without evidence.