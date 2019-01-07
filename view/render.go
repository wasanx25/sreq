package view

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// Render has API URL and API body.
type Render struct {
	URL    string
	Reader io.Reader
}

/*
Item is base on parsing Qiita API JSON.
DOCS: https://qiita.com/api/v2/docs#get-apiv2itemsitem_id
*/
type Item struct {
	Markdown string `json:"body"`
	Title    string `json:"title"`
	URL      string `json:"url"`
}

/*
NewRender creates Render.
URL is Qiita item API URL for getting content.
*/
func NewRender(url string) *Render {
	return &Render{
		URL: url,
	}
}

// GetPage set response body.
func (r *Render) GetPage() error {
	res, err := http.Get(r.URL)
	if err != nil {
		res.Body.Close()
		return err
	}

	r.Reader = res.Body
	return nil
}

// Parse parses Qiita JSON.
func (r *Render) Parse() (item *Item, err error) {
	b, err := ioutil.ReadAll(r.Reader)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(b, &item)
	return
}

func (r *Render) Write(file, body string) (err error) {
	err = ioutil.WriteFile(file, []byte(body), os.ModePerm)
	return
}
