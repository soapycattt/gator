package config

import (
	"testing"
)

func TestRead(t *testing.T) {
	cfg, err := Read()
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	if cfg.DBURL == "" {
		t.Errorf("expected DBURL to not be empty, got %q", cfg.DBURL)
	}
}
