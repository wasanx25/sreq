package view_test

import (
	"bytes"
	"testing"

	"github.com/wasanx25/sreq/view"
)

func TestGetPage(t *testing.T) {
	r := view.NewRender("https://qiita.com")

	err := r.GetPage()
	if err != nil {
		t.Error(err)
	}
}

func TestParse(t *testing.T) {
	r := view.NewRender("https://qiita.com")

	stdin := []byte(`{"rendered_body":"render_body","body":"body","title":"title","url":"url"}`)

	r.Reader = bytes.NewBuffer(stdin)

	item, err := r.Parse()
	if err != nil {
		t.Error(err)
	}

	if item.HTML != "render_body" {
		t.Errorf("expected=%q, got=%q", "render_body", item.HTML)
	}

	if item.Markdown != "body" {
		t.Errorf("expected=%q, got=%q", "body", item.Markdown)
	}

	if item.Title != "title" {
		t.Errorf("expected=%q, got=%q", "title", item.Title)
	}

	if item.URL != "url" {
		t.Errorf("expected=%q, got=%q", "url", item.URL)
	}
}
