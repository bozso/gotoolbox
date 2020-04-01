package table

import (
    "fmt"
)

type Writer interface {
    AddHeader(h Header)
    AddRow(r Row)
    fmt.Stringer
}
