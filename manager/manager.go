package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wasanx25/sreq/history"
	"github.com/wasanx25/sreq/search"
	"github.com/wasanx25/sreq/view"

	"github.com/AlecAivazis/survey"
)

type Manager struct {
	history *history.History
	render *view.Render
	search *search.Search
}

func New(keyword, sort string) *Manager {
	historyFile := filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")

	s := search.New(keyword, sort)
	h := history.New(historyFile)

	return &Manager{
		history: h,
		search: s,
	}
}

func (m *Manager) Run() (err error) {
again:
	page := m.search.GetURL()
	err = m.search.Exec(page)
	if err != nil {
		return
	}

	contents := m.search.Contents
	var contentsStr []string

	for _, c := range contents {
		contentsStr = append(contentsStr, c.Title)
	}
	contentsStr = append(contentsStr, "next")

	contQs := []*survey.Question{
		{
			Name: "content",
			Prompt: &survey.Select{
				Message: "Choice a content",
				Options: contentsStr,
			},
			Validate: survey.Required,
		},
	}

	answer := struct {
		Content string
	}{}

	err = survey.Ask(contQs, &answer)

	if err != nil {
		return
	}

	switch answer.Content {
	case "next":
		goto again
	}
}
