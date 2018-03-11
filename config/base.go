package config

// Qiita based on Qiita API
type Qiita struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Body  string `json:"body"`
}

// BaseURL get qiita api URL
func BaseURL(pagenation string, arg string) string {
	return "http://qiita.com/api/v2/items?page=" + pagenation + "&per_page=10&query=" + arg
}
