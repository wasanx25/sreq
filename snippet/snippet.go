package snippet

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Snippets struct {
	Snippets []SnippetInfo `toml:"snippets"`
}

type SnippetInfo struct {
	SearchKeyword string `toml:"search_keyword"`
	Url           string `toml:"url"`
	Title         string `toml:"title"`
}

func (snippets *Snippets) Load(fileName string) error {
	snippetFile := fileName
	if _, err := toml.DecodeFile(snippetFile, snippets); err != nil {
		return fmt.Errorf("Failed to load snippet file. %v", err)
	}
	return nil
}

func (snippets *Snippets) Save(fileName string) error {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Failed to save snippet file. err: %s", err)
	}
	return toml.NewEncoder(f).Encode(snippets)
}
