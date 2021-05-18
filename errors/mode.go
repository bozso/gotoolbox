package errors

import (
	"fmt"
)

type Mode string

func (m Mode) Error(got string) (err error) {
	return UnrecognizedModeError{
		name: string(m),
		got:  got,
	}
}

func UnrecognizedMode(got, name string) error {
	return UnrecognizedModeError{name, got, nil}
}

type UnrecognizedModeError struct {
	name, got string
	err       error
}

func (e UnrecognizedModeError) Error() string {
	return fmt.Sprintf("unrecognized mode '%s' for %s", e.got, e.name)
}

func (e UnrecognizedModeError) Unwrap() error {
	return e.err
}
