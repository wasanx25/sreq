package src

// Qiita based on Qiita API
type Qiita struct {
	HTML     string `json:"rendered_body"`
	Markdown string `json:"body"`
	Title    string `json:"title"`
	URL      string `json:"url"`
}

// PageURL is that get scraping page url
func PageURL(keywords string, sort string, pagenation string) string {
	return "https://qiita.com/search?page=" + pagenation + "&q=" + keywords + "&sort=" + sort
}

// APIURL is that get api url
func APIURL(id string) string {
	return "https://qiita.com/api/v2/items/" + id
}
