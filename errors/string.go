package errors

import (
    "fmt"
)

type NotEmpty string

func (n NotEmpty) Check(variable string) (err error) {
    if len(n) == 0 {
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
