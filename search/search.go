package search

import (
	"bytes"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

type Search struct {
	Keyword    string
	Pagenation int
	Sort       string
	Contents   []*Content
}

// Content is structure that scraping content from Qiita
type Content struct {
	ID    string
	Title string
	Desc  string
}

func New(keyword string, sort string) *Search {
	return &Search{
		Keyword:    keyword,
		Pagenation: 0,
		Sort:       sort,
	}
}

func (s *Search) GetURL() string {
	q := url.Values{}
	q.Set("page", strconv.Itoa(s.Pagenation))
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

func (s *Search) NextPage() {
	s.Pagenation++
}

func (s *Search) Exec(page string) ([]*Content, error) {
	doc, err := goquery.NewDocument(page)
	if err != nil {
		return nil, err
	}
	doc.Find(".searchResult").Each(s.getAttr)

	return s.Contents, nil
}

func (s *Search) ContentString() string {
	var out bytes.Buffer

	for n, c := range s.Contents {
		out.WriteString(color.YellowString(strconv.Itoa(n) + " -> "))
		out.WriteString(c.Title + "\n")
		out.WriteString(color.GreenString(c.Desc) + "\n\n")
	}

	if len(s.Contents) == 10 {
		out.WriteString(color.YellowString("n -> ") + "next page\n")
	}

	return out.String()
}

func (s *Search) getAttr(_ int, q *goquery.Selection) {
	itemID, _ := q.Attr("data-uuid")
	title := q.Find(".searchResult_itemTitle a").Text()
	desc := q.Find(".searchResult_snippet").Text()

	content := &Content{
		ID:    itemID,
		Title: title,
		Desc:  desc,
	}

	s.Contents = append(s.Contents, content)
}
