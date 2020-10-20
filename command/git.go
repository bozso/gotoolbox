package command

type Git struct {
    git Caller
}

func NewGit(command Caller) (g Git) {
    return Git {
        git: command,
    }
}

func (g Git) Status() (b []byte, err error) {
    return g.git.Call("status")
}
