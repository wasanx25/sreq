package config

import (
	"os"
	"path/filepath"
)

func HistoryFile() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")
}
