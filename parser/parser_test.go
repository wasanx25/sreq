package parser_test

import (
	"net/url"
	"testing"

	"github.com/wasanx25/sreq/parser"

	"gopkg.in/h2non/gock.v1"
)

func TestParse(t *testing.T) {
	defer gock.Off()

	gock.New("https://test.com").
		Get("/test").
		Reply(200).
		JSON(map[string]string{
			"body": "test body (markdown)",
			"title": "test title",
			"url": "http://test.url",
		})

	u := url.URL{
		Scheme: "https",
		Host: "test.com",
		Path: "test",
	}

	item, err := parser.Parse(u)
	if err != nil {
		t.Fatal(err)
	}

	if item.Markdown != "test body (markdown)" {
		t.Errorf("expected: %v, got: %v", "test body (markdown)", item.Markdown)
	}

	if item.Title != "test title" {
		t.Errorf("expected: %v, got: %v", "test title", item.Title)
	}

	if item.URL != "http://test.url" {
		t.Errorf("expected: %v, got: %v", "http://test.url", item.URL)
	}
}