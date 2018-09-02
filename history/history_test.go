package history_test

import (
	"os"
	"testing"

	"github.com/wasanx25/sreq/history"
)

func TestRead(t *testing.T) {
	f, err := os.Create("./test.toml")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	defer os.Remove("./test.toml")

	h := history.New("./test.toml")
	err = h.Read()
	if err != nil {
		t.Error(err)
	}
}

func TestWrite(t *testing.T) {
	f, err := os.Create("./test.toml")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	defer os.Remove("./test.toml")

	h := history.New("./test.toml")

	err = h.Write("testK", "testU", "testT")
	if err != nil {
		t.Error(err)
	}

	actualKeyword := h.Snippets.Snippets[0].SearchKeyword
	actualURL := h.Snippets.Snippets[0].URL
	actualTitle := h.Snippets.Snippets[0].Title

	if actualKeyword != "testK" {
		t.Errorf("expected=%q, got=%q", "testK", actualKeyword)
	}

	if actualURL != "testU" {
		t.Errorf("expected=%q, got=%q", "testU", actualURL)
	}

	if actualTitle != "testT" {
		t.Errorf("expected=%q, got=%q", "testT", actualTitle)
	}
}
