package command

import (
    "os/exec"
)

/*
A simple wrapper around a command, handling an executable binary to be
started and used as process and handled by another struct.
*/
type Process struct {
    process *exec.Cmd
}

func NewProcess(exec *exec.Cmd) (c Command) {
    return Command {
        matlab: exec,
    }
}


/*
A wrapper around a process that pipes input to the started process,
handles output and errors returned by the process.
*/
type StartedProcess struct {
    Process
    Stderr io.ReadCloser
    Stdout io.ReadCloser
    Stdin  io.WriteCloser
}

/*
Start a process and setup it's pipes.
 */
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

/* 
Send data to the started process. After writing to the stdin pipe,
the function reads from the stderr and reports any errors.
 */
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

// Waits for the process to close.
func (s StartedProcess) Wait() (err error) {
    return s.process.Wait()
}

/*
An error describing the context of a failed communication with a
process.
 */
type Error struct {
    Exec string,
    Message []byte,
    err error
}

// Implements the error interface.
func (e Error) Error() (s string) {
    return fmt.Sprintf("process '%s' returned with error: '%s'",
        e.Exec, e.Message)
}

// Can be unwrapped.
func (e Error) Unwrap() (err error) {
    return e.err
}
