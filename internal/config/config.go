package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	HTTPAddr        string
	GRPCAddr        string
	DatabaseURL     string
	ShutdownTimeout time.Duration
}

func Load() (*Config, error) {
	httpAddr := getEnv("HTTP_ADDR", ":8080")
	grpcAddr := getEnv("GRPC_ADDR", ":9090")
	dbURL := getEnv("DATABASE_URL", "postgres://app:app@localhost:5432/app?sslmode=disable")
	shutdownTimeout := getEnvDuration("SHUTDOWN_TIMEOUT", 5*time.Second)

	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is empty")
	}

	return &Config{
		HTTPAddr:        httpAddr,
		GRPCAddr:        grpcAddr,
		DatabaseURL:     dbURL,
		ShutdownTimeout: shutdownTimeout,
	}, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
