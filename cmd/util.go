package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func openLynx(html string) {
	execCmd(html, "lynx", "/tmp/sreq.html")
}

func openEditor(body string, editor string) {
	execCmd(body, editor, "/tmp/sreq.txt")
}

func execCmd(body string, cmdName string, file string) {
	text := []byte(body)
	ioutil.WriteFile(file, text, os.ModePerm)
	cmd := exec.Command(cmdName, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
