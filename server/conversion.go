package server

import (
    "fmt"
)

type ConvertFail struct {
    name fmt.Stringer
}

func NewConvertFail(name fmt.Stringer) (c ConvertFail) {
    c.name = name
    return
}

func GenerateFails(names ...fmt.Stringer) (fails []ConvertFail) {
    fails = make([]ConvertFail, 0, len(names))
    
    for ii, name := range names {
        fails[ii] = NewConvertFail(name)
    }
    return
}

func (c ConvertFail) Error() (s string) {
    return fmt.Sprintf("failed to convert database entity to %s",
        c.name.String())
}
