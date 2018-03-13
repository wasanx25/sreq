package config

import (
	"os"
	"path/filepath"
)

// HistoryFile get command-history file-name
func HistoryFile() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")
}
