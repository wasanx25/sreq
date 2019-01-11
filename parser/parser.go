package parser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
Item is base on parsing Qiita API JSON.
DOCS: https://qiita.com/api/v2/docs#get-apiv2itemsitem_id
*/
type item struct {
	Markdown string `json:"body"`
	Title    string `json:"title"`
	URL      string `json:"url"`
}

func Parse(u url.URL) (*item, error) {
	res, err := http.Get(u.String())
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	i := &item{}

	err = json.Unmarshal(b, i)
	return i, err
}
