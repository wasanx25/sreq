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
	expectedURL := "https://qiita.com/search?pagenation=0&q=testK&sort=testS"
	actual := s.GetURL()

	if actual != expectedURL {
		t.Errorf("expected=%q, got=%q", expectedURL, actual)
	}
}

func TestExec(t *testing.T) {
	s := search.New("testK", "testS")
	t.Run("return content", func(t *testing.T) {
		actualC, actualE := s.Exec()
		expectedContents := []*search.Content{}
		var expectedError *error

		if actualC != expectedContents {
			t.Errorf("expected=%q, got=%q", expectedContents, actualC)
		}

		if actualE != expectedError {
			t.Errorf("expected=%q, got=%q", expectedError, actualE)
		}
	})
}
