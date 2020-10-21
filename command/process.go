package command

import (
    "os/exec"
)

type Process struct {
    process *exec.Cmd
}

func NewProcess(exec *exec.Cmd) (c Command) {
    return Command {
        matlab: exec,
    }
}

type StartedProcess struct {
    Process
    Stderr io.ReadCloser
    Stdout io.ReadCloser
    Stdin  io.WriteCloser
}

func (p Process) Start() (s StartedProcess, err error) {
    if s.Stderr, err = p.process.StderrPipe(); err != nil {
        return
    }

    if s.Stdout, err = p.process.StdoutPipe(); err != nil {
        return
    }

    if s.Stdin, err = p.process.StdinPipe(); err != nil {
        return
    }

    if err = p.process.Start(); err != nil {
        return
    }

    s.Process = p
    return
}

func (s *StartedProcess) Write(b []byte) (n int, err error) {
    if n, err = s.Stdin.Write(b); err != nil {
        return
    }
    
    out, err := ioutil.ReadAll(s.Stderr)
    if err != nil {
        err = Error{
            Exec: s.process.matlab.String()
            Message: out,
            err: err,
        }
        return
    }
}

func (s StartedProcess) Wait() (err error) {
    return s.process.Wait()
}

type Error struct {
    Exec string,
    Message []byte,
    err error
}

func (e Error) Error() (s string) {
    return fmt.Sprintf("process '%s' returned with error: '%s'",
        e.Exec, e.Message)
}

func (e Error) Unwrap() (err error) {
    return e.err
}
