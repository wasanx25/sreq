package snippet

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Snippets array of SnippetInfo
type Snippets struct {
	Snippets []SnippetInfo `toml:"snippets"`
}

// SnippetInfo for search of qiita
type SnippetInfo struct {
	SearchKeyword string `toml:"search_keyword"`
	Url           string `toml:"url"`
	Title         string `toml:"title"`
}

// Load reading snippet file
func (snippets *Snippets) Load(fileName string) error {
	snippetFile := fileName
	if _, err := toml.DecodeFile(snippetFile, snippets); err != nil {
		return fmt.Errorf("Failed to load snippet file. %v", err)
	}
	return nil
}

// Save snippet file
func (snippets *Snippets) Save(fileName string) error {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Failed to save snippet file. err: %s", err)
	}
	return toml.NewEncoder(f).Encode(snippets)
}
