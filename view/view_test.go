package view_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/wasanx25/sreq/view"
)

func TestExec(t *testing.T) {
	data := []byte("testContent")
	err := ioutil.WriteFile("./test", data, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer os.Remove("./test")

	v := view.New(view.LESS, "./test")

	out := captureStdout(func() {
		v.Exec()
	})

	if out != "testContent" {
		t.Errorf("expected=%q, got=%q", "testContent", out)
	}
}

func captureStdout(fn func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	fn()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
