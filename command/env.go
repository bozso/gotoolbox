package command

import (
    "fmt"
    "os/exec"

    "github.com/bozso/gotoolbox/environment"
)

func AddEnv(e environment.Env) (a AddEnvironment) {
    return AddEnvironment{e}
}

func SetEnv(e environment.Env) (s SetEnvironment) {
    return SetEnvironment{e}
}

type SetEnvironment struct {
    env environment.Env
}

func (s SetEnvironment) ConfigureCommand(cmd *exec.Cmd) {
    cmd.Env = s.env.Get()
}

type AddEnvironment struct {
    env environment.Env
}

func (a AddEnvironment) ConfigureCommand(cmd *exec.Cmd) {
    cmd.Env = append(cmd.Env, a.env.Get()...)
}

