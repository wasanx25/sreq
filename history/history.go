package history

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type History struct {
	File     string
	Snippets *snippets
}

func New(file string) *History {
	s := &snippets{}
	return &History{
		File:     file,
		Snippets: s,
	}
}

func (h *History) Write(keyword, url, title string) (err error) {
	err = h.Snippets.load(h.File)
	if err != nil {
		return err
	}

	s := snippet{
		SearchKeyword: keyword,
		URL:           url,
		Title:         title,
	}

	h.Snippets.Snippets = append(h.Snippets.Snippets, s)
	err = h.Snippets.save(h.File)
	if err != nil {
		return err
	}

	return
}

type snippets struct {
	Snippets []snippet `toml:"snippets"`
}

type snippet struct {
	SearchKeyword string `toml:"search_keyword"`
	URL           string `toml:"url"`
	Title         string `toml:"title"`
}

func (snippets *snippets) load(file string) error {
	if _, err := toml.DecodeFile(file, snippets); err != nil {
		return fmt.Errorf("Failed to load snippet file. %v", err)
	}
	return nil
}

func (snippets *snippets) save(file string) error {
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Failed to save snippet file. err: %s", err)
	}
	return toml.NewEncoder(f).Encode(snippets)
}
