package control

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/wasanx25/sreq/history"
	"github.com/wasanx25/sreq/search"
	"github.com/wasanx25/sreq/view"
	"github.com/wasanx25/goss"
)

/*
Control has search option and history and viewing.
User control on CLI.
*/
type Control struct {
	Num     int
	Next    bool
	History *history.History
	Render  *view.Render
	Search  *search.Search
	View    *view.View
}

/*
New creates Control structure.
Set history and result filename.
*/
func New(keyword, sort string) *Control {
	hFile := filepath.Join(os.Getenv("HOME"), ".config", "sreq", "sreq-history.toml")

	var (
		cmd     view.Command
		file    string
		options []string
	)

	s := search.New(keyword, sort)
	h := history.New(hFile)
	v := view.New(cmd, file, options)

	return &Control{
		Search:  s,
		History: h,
		View:    v,
	}
}

// Exec is main method in search command.
func (c *Control) Exec() (err error) {
loop:
	for {
		page := c.Search.GetURL()
		err = c.Search.Exec(page)
		if err != nil {
			break loop
		}

		fmt.Println(c.Search.ContentString())
		err = c.Scan()
		if err != nil {
			break loop
		}
		if !c.Next {
			content := c.Search.Contents[c.Num]
			url := c.GetURL(content.ID)
			c.Render = view.NewRender(url)
			if err = c.Render.GetPage(); err != nil {
				break loop
			}

			item, err := c.Render.Parse()
			if err != nil {
				break loop
			}

			goss.Run(item.Markdown)

			c.History.Write(c.Search.Keyword, item.URL, item.Title)
			c.View.Exec()
			break loop
		}
	}

	return
}

// Scan execute fmt.Scanf and get content index.
func (c *Control) Scan() (err error) {
	var num string
	if _, err = fmt.Scanf("%s", &num); err != nil {
		return
	}

	if num == "n" {
		c.Next = true
		c.Search.NextPage()
	} else {
		c.Next = false
		i, _ := strconv.Atoi(num)
		c.Num = i
	}

	return
}

// GetURL get item API URL.
func (c *Control) GetURL(id string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "qiita.com",
		Path:   filepath.Join("api", "v2", "items", id),
	}
	return u.String()
}
