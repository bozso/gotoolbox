package command

import (
    "context"

    "os/exec"
)

type Creator interface {
    Create(name string, args ...string) *exec.Cmd
}

type DefaultCreator struct {}

func (_ DefaultCreator) Create(name string, args ...string) (cmd *exec.Cmd) {
    return exec.Command(name, args...)
}

type CreateWithContext struct {
    ctx context.Context
}

func (c CreateWithContext) CreateCmd(name string, args ...string) (cmd *exec.Cmd) {
    return exec.CommandContext(c.ctx, name, args...)
}

func WithContext(ctx context.Context) (c CreateWithContext) {
    return CreateWithContext{ctx}
}

type Configure interface {
    ConfigureCommand(*exec.Cmd)
}

type Configures struct {
    configs []Configure
}

func (c Configures) ConfigureCommand(cmd *exec.Cmd) {
    for _, conf := range c.configs {
        conf.ConfigureCommand(cmd)
    }
}


type FrozenArguments struct {
    creator Creator
}
