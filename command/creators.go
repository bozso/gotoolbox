package command

import (
    "io"
    "os/exec"
)


type PathChecked struct {}

var PathChecker = PathChecked{}

func (_ PathChecked) CreateCmd(name string, args ...string) (c *exec.Cmd, err error) {
    // TODO: implement
    return
}


type DebugBuilder struct {
    writer io.Writer
}

func NewDebugBuilder(w io.Writer) (d DebugBuilder) {
    d.writer = w
    return
}

func (d DebugBuilder) NewCommand(execPath string) (c Command, err error) {
    c = New(NewDebug(d.writer, execPath))
    return
}
