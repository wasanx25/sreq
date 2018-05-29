package src

import (
	"fmt"
)

// Searcher is very nice
type Searcher struct {
	Keywords   string
	Pagination int
	Sort       string
	Lynx       bool
}

// Exec is scraping and viewing contents and selecting contents
func (s *Searcher) Exec() {
	for {
		contents, err := search(s.Keywords, s.Pagination, s.Sort)
		if err != nil {
			fmt.Println(err)
			break
		}
		viewList(contents)
		endPhase := scan(contents, s.Keywords, s.Lynx)
		if endPhase {
			break
		}
		s.Pagination++
	}
}
