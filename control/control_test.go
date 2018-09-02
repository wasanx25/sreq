package control_test

import (
	"testing"

	"github.com/wasanx25/sreq/control"
)

func TestGetURL(t *testing.T) {
	c := &control.Control{}

	actual := c.GetURL("test")
	expected := "https://qiita.com/api/v2/items/test"

	if actual != expected {
		t.Errorf("expected=%q, got=%q", expected, actual)
	}
}
