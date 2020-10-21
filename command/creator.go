package command

import (
    "os/exec"
)

type Creator interface {
    Create() *exec.Cmd
}

type FrozenArguments struct {
    creator Creator
}
