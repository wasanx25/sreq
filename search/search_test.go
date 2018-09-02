package search_test

import (
	"strconv"
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
	// TODO: use dummy URL that running TCP
	actualError := s.Exec("https://qiita.com")
	var expectedContents []*search.Content
	var expectedError error

	for i, content := range s.Contents {
		if content != expectedContents[i] {
			t.Errorf("expected=%q, got=%q", expectedContents[i], content)
		}
	}

	if actualError != expectedError {
		t.Errorf("expected=%q, got=%q", expectedError, actualError)
	}
}

func TestContentString(t *testing.T) {
	tests := []struct {
		contentCount int
		expected     string
	}{
		{
			2,
			`0 -> title_1
desc_1

1 -> title_2
desc_2

`,
		},
		{
			10,
			`0 -> title_1
desc_1

1 -> title_2
desc_2

2 -> title_3
desc_3

3 -> title_4
desc_4

4 -> title_5
desc_5

5 -> title_6
desc_6

6 -> title_7
desc_7

7 -> title_8
desc_8

8 -> title_9
desc_9

9 -> title_10
desc_10

n -> next page
`,
		},
	}

	s := search.New("testK", "testS")

	for _, tt := range tests {
		num := 1
		for {
			id := strconv.Itoa(num)
			title := "title_" + id
			desc := "desc_" + id
			s.Contents = append(
				s.Contents,
				&search.Content{ID: id, Title: title, Desc: desc},
			)
			if num == tt.contentCount {
				break
			}
			num++
		}

		actual := s.ContentString()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}

		s.Contents = nil
	}
}
