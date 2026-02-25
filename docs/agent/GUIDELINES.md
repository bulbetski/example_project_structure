# Agent Operating Rules

## General Principles
- Do not write production code without a plan.
- Use `docs/agent/AGENTS.md` and `docs/` as the Memory Bank before making decisions.
- If documentation conflicts with code, report it before proceeding.
- If something is unclear, mark it `UNKNOWN` instead of guessing.

## Workflow (Mandatory)

### 1) Context
Before implementing anything:
- Re-read `docs/agent/AGENTS.md`.
- Re-read relevant docs in `docs/`.
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
- Before commit, run `make run` and verify HTTP health with `curl -i http://localhost:8080/healthz` (expect `200`).
- If environment restrictions prevent live run checks, run `GOCACHE=/tmp/go-build-cache go test ./...` and report the limitation.

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

## Memory Bank
- `docs/agent/AGENTS.md` is the project constitution.
- `docs/` contains structural truth.
- If repeated mistakes occur, update `docs/agent/AGENTS.md`.
- If new architectural decisions are made, update `docs/`.

## Forbidden
- Writing code without plan.
- Skipping tests.
- Ignoring `docs/agent/AGENTS.md`.
- Modifying unrelated modules.
- Making assumptions without evidence.
