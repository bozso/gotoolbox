package command

import (
    "context"

    "os/exec"
)

type Creator interface {
    Create(name string, args ...string) (*exec.Cmd, error)
}

type DefaultCreator struct {}

func (_ DefaultCreator) NewCmd(name string, args ...string) (c *exec.Cmd, err error) {
    return exec.Command(name, args...), nil
}

type CreateWithContext struct {
    ctx context.Context
}

func (cw CreateWithContext) NewCmd(name string, args ...string) (c *exec.Cmd, err error) {
    return exec.CommandContext(cw.ctx, name, args...), nil
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
