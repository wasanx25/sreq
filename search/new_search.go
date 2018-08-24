package search

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
