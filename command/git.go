package command

import (

)

/*
A wrapper around the git command line application. Implements
the VCS interface.
*/
type Git struct {
    git Caller
}

func NewGit(command Caller) (g Git) {
    return Git {
        git: command,
    }
}

func (g Git) Status() (b []byte, err error) {
    return g.git.Call([]string{"status"})
}
