package config

import (
	"path/filepath"
	"testing"
)

func TestNewConfig(t *testing.T) {
	configPath := filepath.Join("..", "config.toml")

	config := NewConfig(configPath)

	if config.Server.Port != ":5000" {
		t.Errorf("expected 5000, but got %s", config.Server.Port)
	}

	if config.Grpc.Port != ":5001" {
		t.Errorf("expected 5051, but got %s", config.Grpc.Port)
	}
}
