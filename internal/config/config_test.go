package config

import "testing"

func TestFromEnvironmentRejectsInvalidShutdownTimeout(t *testing.T) {
	t.Setenv("CSG_SHUTDOWN_TIMEOUT", "eventually")

	if _, err := FromEnvironment(); err == nil {
		t.Fatal("expected invalid shutdown timeout to be rejected")
	}
}

func TestFromEnvironmentUsesExplicitAddress(t *testing.T) {
	t.Setenv("CSG_HTTP_ADDRESS", "127.0.0.1:9090")

	cfg, err := FromEnvironment()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.HTTPAddress != "127.0.0.1:9090" {
		t.Fatalf("unexpected address: %s", cfg.HTTPAddress)
	}
}
