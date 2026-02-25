COMPOSE_FILE := build/local/docker-compose.yml
GOCACHE_DIR := /tmp/go-build-cache

.PHONY: deps run stop test test-cov codegen-deps codegen-update

deps:
	@command -v go >/dev/null || { echo "go is required: https://go.dev/dl/"; exit 1; }
	@command -v docker >/dev/null || { echo "docker is required: https://docs.docker.com/get-docker/"; exit 1; }
	@docker compose version >/dev/null 2>&1 || { echo "docker compose is required"; exit 1; }
	@docker info >/dev/null 2>&1 || { echo "docker daemon is not available; start Docker and retry"; exit 1; }
	go mod download

run: deps
	docker compose -f $(COMPOSE_FILE) up -d db
	@db_cid=$$(docker compose -f $(COMPOSE_FILE) ps -q db); \
	if [ -z "$$db_cid" ]; then echo "db container was not created"; exit 1; fi; \
	for i in $$(seq 1 30); do \
		status=$$(docker inspect -f '{{if .State.Health}}{{.State.Health.Status}}{{else}}{{.State.Status}}{{end}}' $$db_cid 2>/dev/null || true); \
		if [ "$$status" = "healthy" ]; then break; fi; \
		if [ "$$i" -eq 30 ]; then \
			echo "db did not become healthy in time"; \
			docker compose -f $(COMPOSE_FILE) logs db; \
			exit 1; \
		fi; \
		sleep 1; \
	done
	go run ./cmd/app

stop:
	docker compose -f $(COMPOSE_FILE) down

test:
	GOCACHE=$(GOCACHE_DIR) go test ./... -v --race --tags=tests

test-cov:
	GOCACHE=$(GOCACHE_DIR) go test ./... -coverprofile=coverage.txt -covermode=atomic --tags=tests
	@go tool cover -func=coverage.txt | tail -n 1
	@rm -f coverage.txt

codegen-update:
	@./scripts/system/proto-gen.sh

codegen-deps:
	@./scripts/system/proto-deps.sh
