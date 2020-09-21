package command

import (

)

type Builder interface {
    New(string) (Command, error)
}

type SimpleBuilder struct {}

func (s SimpleBuilder) New(execPath string) (c Command, err error)  {
    c = New(execPath)
    return
}

type PathCheckedBuilder struct {
    builder Builder
}

func NewPathCheckedBuilder(b Builder) (p PathCheckedBuilder) {
    p.builder = b
    return
}

func (p PathCheckedBuilder) New(execPath string) (c Command, err error) {
    vf, err := path.New(cmd).ToValidFile()
    
    if err != nil {
        return
    }
    
    c, err = p.builder.New(vf.String)
    return
}

type DebugBuilder struct {
    builder Builder
}

func NewPathDebugBuilder(b Builder) (d DebugBuilder) {
    d.builder = b
    return
}

func (d DebugBuilder) New(execPath string) (c Command, err error) {
    c = d.builder.New(execPath).Debug()
    return
}
