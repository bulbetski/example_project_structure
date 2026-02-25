FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd/app

FROM alpine:3.19

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=build /bin/app /app/app

ENV HTTP_ADDR=:8080 \
    GRPC_ADDR=:9090

EXPOSE 8080 9090

USER app

ENTRYPOINT ["/app/app"]
