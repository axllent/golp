package app

import (
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	QuietLogging = true
	Conf.ConfigFile = filepath.Join("..", "test-data", "golp.yaml")

	if err := ParseConfig(); err != nil {
		t.Error(err)
	}
}

func TestConfigFailNoProcesses(t *testing.T) {
	QuietLogging = true
	Conf.ConfigFile = filepath.Join("..", "test-data", "golp-no-processes.yaml")

	if err := ParseConfig(); err == nil {
		t.Error("golp-no-processes.yaml should have returned an error")
	}
}

func TestConfigFailNoDist(t *testing.T) {
	QuietLogging = true
	Conf.ConfigFile = filepath.Join("..", "test-data", "golp-no-dist.yaml")

	if err := ParseConfig(); err == nil {
		t.Error("golp-no-dist.yaml should have returned an error")
	}
}

func TestConfigFailNoSrc(t *testing.T) {
	QuietLogging = true
	Conf.ConfigFile = filepath.Join("..", "test-data", "golp-no-src.yaml")

	if err := ParseConfig(); err == nil {
		t.Error("golp-no-src.yaml should have returned an error")
	}
}

func TestConfigFailNoFile(t *testing.T) {
	QuietLogging = true
	Conf.ConfigFile = filepath.Join("..", "test-data", "golp-does-not-exist.yaml")

	if err := ParseConfig(); err == nil {
		t.Error("golp-does-not-exist.yaml should have returned an error")
	}
}
