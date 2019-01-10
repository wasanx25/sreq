package pager_test

import (
	"github.com/wasanx25/sreq/pager"
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	} {
		{"rel", true},
		{"created", true},
		{"stock", true},
		{"none", false},
	}

	for _, tt := range tests {
		s := pager.SortType(tt.input)
		actual := s.Valid()
		if actual != tt.expected {
			t.Errorf("Valid() not expected(%v), got(%v) ", tt.expected, actual)
		}
	}
}

func TestGetURL(t *testing.T) {
	expected := "https://qiita.com/search?page=0&q=keyword&sort=rel"
	p, err := pager.New("keyword", "rel")

	if err != nil {
		t.Fatalf("occur err message: (%v) ", err)
	}

	url := p.GetURL()

	if url != expected {
		t.Errorf("GetURL() not expected(%v), got(%v) ", expected, url)
	}
}