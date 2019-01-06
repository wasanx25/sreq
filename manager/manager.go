package manager

import (
	"github.com/wasanx25/goss"
	"net/url"
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

	var content *search.Content
	switch answer.Content {
	case "next":
		goto again
	default:
		for _, c := range contents {
			if answer.Content == c.Title {
				content = c
				break
			}
		}
	}

	u := url.URL{
		Scheme: "https",
		Host:   "qiita.com",
		Path:   filepath.Join("api", "v2", "items", content.ID),
	}

	m.render = view.NewRender(u.String())
	err = m.render.GetPage()
	if err != nil {
		return
	}

	item, err := m.render.Parse()
	if err != nil {
		return
	}

	err = goss.Run(item.Markdown)
	if err != nil {
		return
	}

	m.history.Write(m.search.Keyword, item.URL, item.Title)

	return
}
