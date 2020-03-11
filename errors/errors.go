package errors

import "fmt"

func Wrap(err1, err2 error) error {
    return fmt.Errorf("%w\n%w", err1, err2)
}

func WrapFmt(err error, msg string, args ...interface{}) error {
    s := fmt.Sprintf(msg, args...)
    
    return fmt.Errorf("%s\n%w", s, err)
}

type ErrorBase struct {
    err error
}

func (e *ErrorBase) Wrap(err error) *ErrorBase {
    e.err = err
    return e
}

func (e ErrorBase) Unwrap() error {
    return e.err
}

func NotEmpty(variable, s string) (err error) {
    if len(s) == 0 {
        return EmptyStringError{variable, nil}
    }
    
    return nil
}

type EmptyStringError struct {
    variable string
    err      error
}

func (e EmptyStringError) Error() (s string) {
    s = "expected non empty string"
    
    if v := e.variable; len(v) > 0 {
        s = fmt.Sprintf("%s for '%s'", s, v)
    }
    
    return
}

func (e EmptyStringError) Unwrap() error {
    return e.err
}

func UnrecognizedMode(got, name string) error {
    return UnrecognizedModeError{name, got, nil}
}

type UnrecognizedModeError struct {
    name, got string
    err error
}

func (e UnrecognizedModeError) Error() string {
    return fmt.Sprintf("unrecognized mode '%s' for %s", e.got, e.name)
}

func (e UnrecognizedModeError) Unwrap() error {
    return e.err
}

type OutOfBoundError struct {
    idx, length int
    err error
}

func (o OutOfBoundError) Error() string {
    return fmt.Sprintf("idx '%d' is out of bounds of length '%d'",
        o.idx, o.length)
}

func (o OutOfBoundError) Unwrap() error {
    return o.err
}

func IsOutOfBounds(idx, length int) error {
    if idx >= length {
        return OutOfBoundError{idx:idx, length:length}
    }
    return nil
}
