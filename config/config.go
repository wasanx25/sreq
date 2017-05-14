package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func HistoryFile() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")
}

func KeywordFile() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-keywords.toml")
}

func ConfigFile() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "sreq", "config.toml")
}

type Config struct {
	General GeneralConfig
}

type GeneralConfig struct {
	OutputType string `toml:"output_type"`
	Editor     string `toml:"editor"`
}

func (cfg *Config) Load() error {
	configFile := ConfigFile()
	if _, err := toml.DecodeFile(configFile, cfg); err != nil {
		return fmt.Errorf("Failed to load snippet file. %v", err)
	}
	return nil
}

func (cfg *Config) Save() error {
	configFile := ConfigFile()
	f, err := os.Create(configFile)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Failed to save snippet file. err: %s", err)
	}
	return toml.NewEncoder(f).Encode(cfg)
}
