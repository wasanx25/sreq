package manager

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/wasanx25/sreq/history"
	"github.com/wasanx25/sreq/fetcher"
	"github.com/wasanx25/sreq/pager"
	"github.com/wasanx25/sreq/parser"

	"github.com/AlecAivazis/survey"
	"github.com/wasanx25/goss"
)

type Manager struct {
	history *history.History
	pager   pager.Pager
}

func New(keyword, sort string) *Manager {
	historyFile := filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")

	h := history.New(historyFile)
	// TODO
	p, _ := pager.New(keyword, sort)

	return &Manager{
		history: h,
		pager:   p,
	}
}

func (m *Manager) Run() error {
	var content fetcher.Content

loop:
	for {
		contents, err := fetcher.Fetch(m.pager.GetURL())
		if err != nil {
			return err
		}

		var contentsStr []string

		for _, c := range contents {
			contentsStr = append(contentsStr, c.GetTitle())
		}

		answer, err := m.viewSelector(contentsStr)
		content = m.getContent(answer, contents)

		if content != nil {
			break loop
		}

		m.pager.NextPage()
	}

	u := url.URL{
		Scheme: "https",
		Host:   "qiita.com",
		Path:   filepath.Join("api", "v2", "items", content.GetID()),
	}
	item, err := parser.Parse(u)

	if err != nil {
		return err
	}

	m.history.Write("keyword", item.URL, item.Title)

	err = goss.Run(item.Markdown)

	return err
}

func (m *Manager) viewSelector(contentsStr []string) (string, error) {
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
	err := survey.Ask(contQs, &answer)

	return answer.Content, err
}

func (m *Manager) getContent(answer string, contents []fetcher.Content) fetcher.Content {
	var content fetcher.Content

	switch answer {
	case "next":
		return nil
	default:
		for _, c := range contents {
			if answer == c.GetTitle() {
				content = c
				break
			}
		}
	}

	return content
}
