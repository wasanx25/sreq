package search_test

import (
	"testing"

	"github.com/wasanx25/sreq/search"
)

func TestNew(t *testing.T) {
	actual := search.New("testK", "testS")
	if actual.Keyword != "testK" {
		t.Errorf("expected=%q, got=%q", "testK", actual.Keyword)
	}

	if actual.Sort != "testS" {
		t.Errorf("expected=%q, got=%q", "testS", actual.Sort)
	}
}

func TestGetURL(t *testing.T) {
	s := search.New("testK", "testS")
	expectedURL := "https://qiita.com/search?page=0&q=testK&sort=testS"
	actual := s.GetURL()

	if actual != expectedURL {
		t.Errorf("expected=%q, got=%q", expectedURL, actual)
	}
}

func TestNextPage(t *testing.T) {
	s := search.New("testK", "testS")
	s.NextPage()

	if s.Pagenation != 1 {
		t.Errorf("expected=%q, got=%q", 1, s.Pagenation)
	}
}

func TestExec(t *testing.T) {
	s := search.New("testK", "testS")
	t.Run("return content", func(t *testing.T) {
		actualC, actualE := s.Exec("https://qiita.com")
		var expectedContents []*search.Content
		var expectedError error

		for i, content := range actualC {
			if content != expectedContents[i] {
				t.Errorf("expected=%q, got=%q", expectedContents[i], content)
			}
		}

		if actualE != expectedError {
			t.Errorf("expected=%q, got=%q", expectedError, actualE)
		}
	})
}
