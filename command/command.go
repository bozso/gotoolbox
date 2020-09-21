package command

import (
    "fmt"
    "os/exec"
    "strings"
)

type Command struct {
    cmd string
    debug bool
}

func New(cmd string) (c Command) {
    c.cmd = cmd
    return
}

func (c Command) Debug() (cc Command) {
    cc.cmd, c.debug = c.cmd, true
    return
}

func (c Command) Call(args ...interface{}) (s string, err error) {
    arg := make([]string, len(args))

    for ii, elem := range args {
        arg[ii] = fmt.Sprint(elem)
    }
    
    return c.CallWithArgs(arg...)
}

func (c Command) CallWithArgs(args ...string) (s string, err error) {
    if c.debug {
        fmt.Printf("Debug mode: command to run: '%s %s'\n", c.cmd,
            strings.Join(args, " "))
        return
    }
    
    cmd := exec.Command(c.cmd, args...)
    b, err := cmd.CombinedOutput()
    s = string(b)
    
    if err != nil {
        err = Fail{cmd:c.cmd, out:s, args:args, err:err}
        return
    }

    return
}

type Fail struct {
    cmd, out string
    args []string
    err error
}

func (f Fail) Error() string {
    const errorMessage = "Command '%s %s' failed! Error: '%s'\nOutput of command is: %v"
    
    return fmt.Sprintf(errorMessage,
        f.cmd, strings.Join(f.args, " "), f.err, f.out)
}

func (f *Fail) Wrap(err error) error {
    f.err = err
    return f
}

func (f Fail) Unwrap() error {
    return f.err
}
