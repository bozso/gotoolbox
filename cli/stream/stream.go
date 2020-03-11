package stream

import (
    "fmt"
)

type Name string

type stream struct {
    name Name
}

func (s stream) String() string {
    return string(s.name)
}

func (s stream) fail(op Operation, err error) error {
    return Fail{s.name, op, err}
}

type Operation int

const (
    Read Operation = iota
    Write
)

func (op Operation) String() (s string) {
    switch op {
    case Read:
        s = "read"
    case Write:
        s = "write"
    }
    return s
}

type Fail struct {
    name Name
    op Operation
    err error
}

func (e Fail) Error() string {
    return fmt.Sprintf("failed to %s to stream '%s'",
        e.op.String(), e.name)
}

func (e Fail) Unwrap() error {
    return e.err
}
