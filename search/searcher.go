package search

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

// Validate is that checking options
func (s Searcher) Validate() (err error) {
	if !s.checkSort() {
		err = fmt.Errorf("%s is unknown value", s.Sort)
		return
	}

	if !s.checkKeywords() {
		err = fmt.Errorf("failed to not argument of search keyword")
		return
	}
	return
}

func (s Searcher) checkKeywords() bool {
	if len(s.Keywords) > 0 {
		return true
	}
	return false
}

func (s Searcher) checkSort() bool {
	switch s.Sort {
	case "rel", "created", "stock":
		return true
	default:
		return false
	}
}
