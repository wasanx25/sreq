package src

// Qiita based on Qiita API
type Qiita struct {
	HTML     string `json:"rendered_body"`
	Markdown string `json:"body"`
	Title    string `json:"title"`
	URL      string `json:"url"`
}

// GetPageURL is that get scraping page url
func GetPageURL(keywords string, sort string, pagenation string) string {
	return "https://qiita.com/search?page=" + pagenation + "&q=" + keywords + "&sort=" + sort
}

// GetAPIURL is that get api url
func GetAPIURL(id string) string {
	return "https://qiita.com/api/v2/items/" + id
}
