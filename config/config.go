package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// HistoryFile get command-history file-name
func HistoryFile() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")
}

// GetConfigFile get config file-name
func GetConfigFile() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "sreq", "config.toml")
}

// Config config
type Config struct {
	General GeneralConfig
}

// GeneralConfig general-config
type GeneralConfig struct {
	OutputType string `toml:"output_type"`
	Editor     string `toml:"editor"`
}

// Load load config-file
func (cfg *Config) Load() error {
	configFile := GetConfigFile()
	if _, err := toml.DecodeFile(configFile, cfg); err != nil {
		return fmt.Errorf("Failed to load snippet file. %v", err)
	}
	return nil
}

// Save save config-file
func (cfg *Config) Save() error {
	configFile := GetConfigFile()
	f, err := os.Create(configFile)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Failed to save snippet file. err: %s", err)
	}
	return toml.NewEncoder(f).Encode(cfg)
}
