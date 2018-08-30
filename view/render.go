package view

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Render struct {
	URL    string
	Reader io.Reader
}

type Item struct {
	HTML     string `json:"rendered_body"`
	Markdown string `json:"body"`
	Title    string `json:"title"`
	URL      string `json:"url"`
}

func NewRender(url string) *Render {
	return &Render{
		URL: url,
	}
}

func (r *Render) GetPage() error {
	res, err := http.Get(r.URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	r.Reader = res.Body
	return nil
}

func (r *Render) Parse() (item *Item, err error) {
	b, err := ioutil.ReadAll(r.Reader)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(b, &item)
	return
}
