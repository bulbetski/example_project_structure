# E2E Scenarios

## Preconditions

```bash
docker compose up --build
```

## Scenario 1: HTTP health check

```bash
curl -i http://localhost:8080/healthz
```

## Scenario 2: gRPC health check

```bash
grpcurl -plaintext -d '{}' localhost:9090 grpc.health.v1.Health/Check
```
