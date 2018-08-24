package search

import (
	"net/url"
	"strconv"
)

type search struct {
	Keyword    string
	Pagenation int
	Sort       string
}

func New(keyword string, sort string) *search {
	return &search{
		Keyword:    keyword,
		Pagenation: 0,
		Sort:       sort,
	}
}

func (s *search) GetURL() string {
	q := url.Values{}
	q.Set("pagenetion", strconv.Itoa(s.Pagenation))
	q.Set("q", s.Keyword)
	q.Set("sort", s.Sort)
	u := url.URL{
		Scheme:   "https",
		Host:     "qiita.com",
		Path:     "search",
		RawQuery: q.Encode(),
	}

	return u.String()
}
