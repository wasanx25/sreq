package history

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Snippets array of Snippet
type Snippets struct {
	Snippets []Snippet `toml:"snippets"`
}

// Snippet for search of qiita
type Snippet struct {
	SearchKeyword string `toml:"search_keyword"`
	URL           string `toml:"url"`
	Title         string `toml:"title"`
}

// Load reading snippet file
func (snippets *Snippets) Load() error {
	if _, err := toml.DecodeFile(getHistoryFile(), snippets); err != nil {
		return fmt.Errorf("Failed to load snippet file. %v", err)
	}
	return nil
}

// Save snippet file
func (snippets *Snippets) Save() error {
	f, err := os.Create(getHistoryFile())
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Failed to save snippet file. err: %s", err)
	}
	return toml.NewEncoder(f).Encode(snippets)
}

func getHistoryFile() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")
}