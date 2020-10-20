package tohtml

import (
    "github.com/bozso/gotoolbox/command"
)

type Git struct {
    git command.Caller
}

func NewGit(command command.Caller) (g Git) {
    return Git {
        git: command,
    }
}

func (g Git) Status() (b []byte, err error) {
    return g.git.Call("status")
}
