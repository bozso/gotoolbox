package command

import (
    "fmt"
    "io"
    "io/ioutil"
    "os/exec"
)

/*
ProcessCommand is a wrapper around a path to a valid executable file. It can
be used to start up a process with a set of arguments.
*/
type ProcessCommand struct {
    command path.ValidFile
}


/*
A wrapper around a process that pipes input to the started process,
handles output and errors returned by the process.
*/
type Process struct {
    command *exec.Cmd
    Stderr io.ReadCloser
    Stdout io.ReadCloser
    Stdin  io.WriteCloser
}

/*
Start a process and setup it's pipes.
 */
func (pc ProcessCommand) Start(args ...string) (p Process, err error) {
    cmd := exec.Command(pc.command, args...)

    if s.Stderr, err = cmd.StderrPipe(); err != nil {
        return
    }

    if s.Stdout, err = cmd.StdoutPipe(); err != nil {
        return
    }

    if s.Stdin, err = cmd.StdinPipe(); err != nil {
        return
    }

    if err = cmd.Start(); err != nil {
        return
    }

    s.command = cmd
    return
}

/* 
Send data to the started process. After writing to the stdin pipe,
the function reads from the stderr and reports any errors.
 */
func (p *Process) Write(b []byte) (n int, err error) {
    if n, err = p.Stdin.Write(b); err != nil {
        return
    }

    out, err := ioutil.ReadAll(p.Stderr)
    if err != nil {
        err = Error{
            Exec: p.process.String(),
            Message: out,
            err: err,
        }
    }
    return
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
    // The string describing the executable.
    Exec string
    // Message written to stderr from the executable.
    Message []byte
    // Wrapped error.
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
