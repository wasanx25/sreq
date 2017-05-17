package config

type Qiita struct {
	Title string `json: "title"`
	Url   string `json: "url"`
	Body  string `json: "body"`
}

func BaseURL(pagenation string, arg string) string {
	return "http://qiita.com/api/v2/items?page=" + pagenation + "&per_page=10&query=" + arg
}
