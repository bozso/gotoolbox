package command

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Caller interface {
	Call(args []string) (b []byte, err error)
}

type Command struct {
	Caller
}

func New(caller Caller) (c Command) {
	c.Caller = caller
	return
}

func (c Command) call(args []fmt.Stringer) (b []byte, err error) {
	s := make([]string, len(args))

	for ii, arg := range args {
		s[ii] = arg.String()
	}

	return c.Caller.Call(s)
}

func (c Command) Call(args ...interface{}) (b []byte, err error) {
	s := make([]string, len(args))

	for ii, arg := range args {
		s[ii] = fmt.Sprint(arg)
	}

	return c.Caller.Call(s)
}

type Executable struct {
	cmd string
}

func NewExecutable(cmd string) (e Executable) {
	e.cmd = cmd
	return
}

func (e Executable) Call(args []string) (b []byte, err error) {
	cmd := exec.Command(e.cmd, args...)
	b, err = cmd.CombinedOutput()

	if err != nil {
		err = Fail{cmd: e.cmd, out: b, args: args, err: err}
		return
	}

	return
}

type Debug struct {
	cmd    string
	writer io.Writer
}

func NewDebug(w io.Writer, execPath string) (d Debug) {
	d.writer, d.cmd = w, execPath
	return
}

func (d Debug) Call(args []string) (_ []byte, err error) {
	_, err = fmt.Fprintf(d.writer,
		"Debug mode: command to run: '%s %s'\n", d.cmd,
		strings.Join(args, " "))
	return
}

type Fail struct {
	cmd  string
	out  []byte
	args []string
	err  error
}

func (f Fail) Error() string {
	const errorMessage = "Command '%s %s' failed! Error: '%s'\nOutput of command is: %s"

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
