package errors

import (
	"fmt"
	"os"
)

type Base struct {
	err error
}

func (e *Base) Wrap(err error) (b *Base) {
	e.err = err
	return e
}

func (e Base) Unwrap() error {
	return e.err
}

func Assert(cond bool, msg string, args ...interface{}) {
	if !cond {
		Panicf(msg, args...)
	}
}

func Panicf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	panic("")
}
