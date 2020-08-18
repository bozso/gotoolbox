package server

import (
    "fmt"
)

type ConvertFailLabel string

func (c ConvertFailLabel) New() (cf ConvertFail) {
    cf.name = string(c)
    return
}

func ConversionFailure(str fmt.Stringer) (cf ConvertFail) {
    cf.name = str.String()
    return
}

type ConvertFail struct {
    name string
}

func (c ConvertFail) Error() (s string) {
    return fmt.Sprintf("failed to convert database entity to %s",
        c.name)
}
