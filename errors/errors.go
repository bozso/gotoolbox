package errors

import (
    "log"
    "fmt"
)

type Asserter string

func (a Asserter) Check(err error) {
    if err != nil {
        log.Panicf("%s\n", string(a))
    }
}

func (a Asserter) Checkf(err error, msg string, v ...interface{}) {
    if err != nil {
        s := fmt.Sprintf(msg, v...)
        log.Panicf("%s: %s\n", string(a), s)
    }
}

func Wrap(err1, err2 error) error {
    return fmt.Errorf("%w\n%w", err1, err2)
}

func WrapFmt(err error, msg string, args ...interface{}) error {
    s := fmt.Sprintf(msg, args...)
    
    return fmt.Errorf("%s\n%w", s, err)
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

type Bound int

func (b Bound) IsOutOfBounds(idx int) error {
    if length := int(b); idx >= length {
        return OutOfBoundError{idx:idx, length:length}
    }
    return nil
}

type String string

func (e String) Wrap(err error) (Err error) {
    return fmt.Errorf("%w\n%w", e, err)
}

func (e String) WrapFmt(err error, args ...interface{}) (Err error) {
    s := fmt.Sprintf(string(e), args...)
    return fmt.Errorf("%s\n%w", s, err)
}

func (e String) Error() (s string) {
    return string(e)
}
