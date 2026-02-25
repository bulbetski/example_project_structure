# example app

Minimal Go service scaffold with HTTP + gRPC health checks.

## Run

```bash
docker compose up --build
```

Service ports:
- HTTP: `localhost:8080`
- gRPC: `localhost:9090`

## HTTP API

Health:

```bash
curl -i http://localhost:8080/healthz
```

## gRPC API

Health:

```bash
grpcurl -plaintext -d '{}' localhost:9090 grpc.health.v1.Health/Check
```

List services:

```bash
grpcurl -plaintext localhost:9090 list
```
