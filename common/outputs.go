package common

import (
	"fmt"
	"io/ioutil"
)

// Output interface
type Output interface {
	Show(output string)
}

// TerminalOutput : output in terminal
type TerminalOutput struct{}

func (out TerminalOutput) Show(output string) {
	fmt.Println(output)
}

// FileOutput : output in file
type FileOutput struct {
	Path string
}

func (fileOutput FileOutput) Show(output string) {
	ioutil.WriteFile(fileOutput.Path, []byte(output), 0644)
}
