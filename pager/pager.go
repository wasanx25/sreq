package pager

import (
	"fmt"
	"net/url"
	"strconv"
)

type SortType string

func (s SortType) Valid() bool {
	if s == "rel" || s == "created" || s == "stock" {
		return true
	}
	return false
}

type Pager interface {
	NextPage()
	GetURL() string
}

type pager struct {
	keyword string
	paging int
	sort SortType
}

func New(keyword, sort string) (Pager, error) {
	s := SortType(sort)

	if !s.Valid() {
		return nil, fmt.Errorf("Unknown Value: (%v), select 'rel' or 'created' or 'stock' ", sort)
	}

	return &pager{
		keyword: keyword,
		paging: 0,
		sort: s,
	}, nil
}

func (p *pager) GetURL() string {
	q := url.Values{}
	q.Set("page", strconv.Itoa(p.paging))
	q.Set("q", p.keyword)
	q.Set("sort", string(p.sort))
	u := url.URL{
		Scheme:   "https",
		Host:     "qiita.com",
		Path:     "search",
		RawQuery: q.Encode(),
	}

	return u.String()
}

func (p *pager) NextPage() {
	p.paging++
}
