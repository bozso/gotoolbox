package command

import (
    "io"

    "github.com/bozso/gotoolbox/path"
)

type Builder interface {
    NewCommand(string) (Command, error)
}

type ExecutableBuilder struct {}

func (_ ExecutableBuilder) NewCommand(execPath string) (c Command, err error)  {
    c = New(NewExecutable(execPath))
    return
}

type PathCheckedBuilder struct {
    builder Builder
}

func NewPathCheckedBuilder(b Builder) (p PathCheckedBuilder) {
    p.builder = b
    return
}

func (p PathCheckedBuilder) NewCommand(execPath string) (c Command, err error) {
    vf, err := path.New(execPath).ToValidFile()

    if err != nil {
        return
    }

    c, err = p.builder.New(vf.String())
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
