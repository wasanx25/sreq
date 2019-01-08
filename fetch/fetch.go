package fetch

import (
	"github.com/PuerkitoBio/goquery"
)

type Content interface {
	GetID() string
	GetTitle() string
	GetDesc() string
}

type content struct {
	id    string
	title string
	desc  string
	Content
}

func (c *content) GetID() string {
	return c.id
}

func (c *content) GetTitle() string {
	return c.title
}

func (c *content) GetDesc() string {
	return c.desc
}

func Fetch(url string) (contents []Content, err error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}

	doc.Find(".searchResult").Each(func(_ int, q *goquery.Selection) {
		itemID, _ := q.Attr("data-uuid")
		title := q.Find(".searchResult_itemTitle a").Text()
		desc := q.Find(".searchResult_snippet").Text()

		content := &content{
			id:    itemID,
			title: title,
			desc:  desc,
		}

		contents = append(contents, content)
	})

	return
}
