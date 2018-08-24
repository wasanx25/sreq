package search_test

import (
	"testing"

	"github.com/wasanx25/sreq/search"
)

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
