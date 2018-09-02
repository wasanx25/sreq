package view

import (
	"os"
	"os/exec"
)

// Command is in const string
type Command string

// View has commands string.
type View struct {
	Cmd    Command
	Option []string
	File   string
}

const (
	// LESS command
	LESS = "less"
	// LYNX command
	LYNX = "lynx"
)

/*
New is View initializer.
Commands of executing.
*/
func New(cmd Command, file string, option []string) *View {
	return &View{
		Cmd:    cmd,
		File:   file,
		Option: option,
	}
}

/*
Exec execute standard input and output.
Use View Command for viewing result content.
*/
func (v *View) Exec() {
	var args []string
	args = append(v.Option, v.File)
	cmd := exec.Command(string(v.Cmd), args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
