package view

import (
	"os"
	"os/exec"
)

type Command string

type View struct {
	Cmd    Command
	Option []string
	File   string
}

const (
	LESS = "less"
	LYNX = "lynx"
)

func New(cmd Command, file string, option ...string) *View {
	return &View{
		Cmd:    cmd,
		File:   file,
		Option: option,
	}
}

func (v *View) Exec() {
	var args []string
	args = append(v.Option, v.File)
	cmd := exec.Command(string(v.Cmd), args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
