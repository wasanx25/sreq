package snippet

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var configFile = filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")

type Snippets struct {
	Snippets []SnippetInfo `toml:"snippets"`
}

type SnippetInfo struct {
	SearchKeyword string `toml:"search_keyword"`
	Url           string `toml:"url"`
	Title         string `toml:"title"`
}

func (snippets *Snippets) Load() error {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "sreq")
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		fmt.Errorf("cannot create directory: %v", err)
	}
	snippetFile := configFile
	if _, errr := toml.DecodeFile(snippetFile, snippets); errr != nil {
		return fmt.Errorf("Failed to load snippet file. %v", errr)
	}
	return nil
}

func (snippets *Snippets) Save() error {
	snippetFile := configFile
	f, err := os.Create(snippetFile)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Failed to save snippet file. err: %s", err)
	}
	return toml.NewEncoder(f).Encode(snippets)
}
