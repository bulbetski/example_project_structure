# Project Structure Template

## Canonical Folder Tree

```text
.
├─ Makefile
├─ README.md
├─ build/
│  ├─ app/
│  │  └─ clickhouse_migrations/
│  │     ├─ common/
│  │     ├─ dev/
│  │     ├─ uat/
│  │     ├─ prod/
│  │     └─ embed.go
│  └─ local/
├─ cmd/
│  └─ app/
│     ├─ internal/
│     └─ main.go
├─ docs/
│  ├─ agent/
│  │  ├─ AGENTS.md
│  │  ├─ GUIDELINES.md
│  │  └─ STRUCTURE.md
│  ├─ openapi/
│  └─ protobuf/
├─ internal/
│  ├─ cli/
│  │  └─ deps/
│  ├─ config/
│  ├─ debugserver/
│  ├─ domain/
│  ├─ httptransport/
│  ├─ kafkatransport/
│  ├─ pkg/
│  ├─ repository/
│  ├─ rpctransport/
│  └─ service/
├─ pkg/
├─ scripts/
│  ├─ system/
│  └─ usr/
└─ tests/
   └─ load/
```

## Layer Responsibilities

- `cmd/app/main.go` is the service entrypoint; subcommands live in `cmd/app/internal`.
- `internal/cli/deps` owns dependency providers and DI container wiring used by subcommands.
- `internal/config` holds configuration constants.
- `internal/domain` contains domain types shared by `service` and `repository`.
- `internal/service` implements service logic and defines interfaces for storage dependencies.
- `internal/repository` implements storage interfaces used by `service`.
- `internal/httptransport` and `internal/rpctransport` provide HTTP and gRPC transport handlers.
- `internal/kafkatransport` provides Kafka consumer handlers.
- `internal/pkg` contains internal-only helpers; `pkg` contains reusable helpers.
- `build/local` keeps runtime configs and docker compose for local dependencies.
- `build/app/clickhouse_migrations` holds ClickHouse migrations split by environment.
- `docs/openapi` and `docs/protobuf` store specs for code generation.
- `docs/agent` stores agent operating rules and structure guidance.
- `scripts/system` and `scripts/usr` provide Make targets.
- `tests/load` hosts load tests.

## Dependency Flow (Expected)

```text
cmd/main
  -> cmd/internal (cobra)
    -> internal/cli/deps (providers + container)
      -> transport (http/rpc/kafka)
        -> service
          -> repository

internal/domain types are imported by service and repository.
internal/config is imported by wiring and runtime setup.
```

## Template: Add A New Module/Service

1. Add domain types in `internal/domain/<module>`.
2. Add a service package in `internal/service/<module>`.
3. Define storage interfaces in the service and add `//go:generate go run go.uber.org/mock/mockgen` directives.
4. Implement storage in `internal/repository/<module>`.
5. Add transport handlers in `internal/httptransport/<module>`, `internal/rpctransport/<module>`, or
   `internal/kafkatransport/<module>`.
6. Wire providers and container registrations under `internal/cli/deps` (update `providers.go`, `container.go`, and the
   relevant `services.go`, `repositories.go`, `servers.go`, `kafka-consumers.go`, `cronjobs.go`, `migrate.go`).
7. Expose new services in `internal/service/services.go` when they are part of the shared service layer.
8. Add tests next to each source file and generate mocks into a `mocks/` subfolder.

## Checklist: Recreate This Architecture

- Add `cmd/<service>/main.go` and `cmd/<service>/internal` for cobra subcommands.
- Create `internal/cli/deps` with provider functions and a DI container.
- Split `internal` into `domain`, `service`, `repository`, and transport packages.
- Add `internal/config` and `internal/pkg` for shared config and internal helpers.
- Add `pkg` for reusable helpers.
- Add `docs/openapi` and `docs/protobuf` for specs.
- Add `docs/agent` with `AGENTS.md`, `GUIDELINES.md`, and `STRUCTURE.md`.
- Add `build/local` for runtime configs and `build/app/clickhouse_migrations` for migrations.
- Add `scripts/system` and `scripts/usr` and a `Makefile` that includes them.
- Add `tests/load` for load tests.
