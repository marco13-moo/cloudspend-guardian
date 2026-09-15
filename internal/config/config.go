// Package config validates environment-derived process configuration.
package config

import (
	"fmt"
	"os"
	"time"
)

const (
	defaultHTTPAddress       = ":8080"
	defaultReadHeaderTimeout = 5 * time.Second
	defaultShutdownTimeout   = 10 * time.Second
)

// Config contains operational settings that may vary between deployments.
type Config struct {
	HTTPAddress       string
	ReadHeaderTimeout time.Duration
	ShutdownTimeout   time.Duration
}

// FromEnvironment constructs and validates process configuration.
func FromEnvironment() (Config, error) {
	cfg := Config{
		HTTPAddress:       valueOrDefault("CSG_HTTP_ADDRESS", defaultHTTPAddress),
		ReadHeaderTimeout: defaultReadHeaderTimeout,
		ShutdownTimeout:   defaultShutdownTimeout,
	}

	if raw := os.Getenv("CSG_SHUTDOWN_TIMEOUT"); raw != "" {
		duration, err := time.ParseDuration(raw)
		if err != nil || duration <= 0 {
			return Config{}, fmt.Errorf("CSG_SHUTDOWN_TIMEOUT must be a positive duration: %q", raw)
		}
		cfg.ShutdownTimeout = duration
	}

	return cfg, nil
}

func valueOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
