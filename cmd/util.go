package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// OpenEditor open your choosing editor.
func OpenEditor(body string, editor string) {
	text := []byte(body)
	ioutil.WriteFile("/tmp/sreq.txt", text, os.ModePerm)
	cmd := exec.Command(editor, "/tmp/sreq.txt")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
