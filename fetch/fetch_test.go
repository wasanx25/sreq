package fetch_test

import (
	"github.com/wasanx25/sreq/fetch"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestFetch(t *testing.T) {
	defer gock.Off()

	gock.New("https://test.com").
		Get("/test").
		Reply(200).
		File("testdata/sample.html")

	contents, err := fetch.Fetch("https://test.com/test")
	if err != nil {
		t.Fatal(err)
	}

	if len(contents) < 1 {
		t.Fatal("error")
	}

	for _, c := range contents {
		if c.GetID() != "test1" {
			t.Fatalf("expected: %v, got: %v", "test1", c.GetID())
		}

		if c.GetTitle() != "test title 1" {
			t.Fatalf("expected: %v, got: %v", "test title 1", c.GetTitle())
		}

		if c.GetDesc() != "test snippet 1" {
			t.Fatalf("expected: %v, got: %v", "test snippet 1", c.GetDesc())
		}
	}
}