package table

import (
    "strings"

    "github.com/bozso/gotoolbox/errors"
)

type Mode int

const (
    Html Mode = iota
    Latex
)

func (m Mode) String() (s string) {
    switch m {
    case Html:
        s = "html"
    case Latex:
        s = "latex"
    }
    
    return
}

func (m *Mode) Set(s string) (err error) {
    sm := strings.ToLower(s)
    
    switch sm {
    case "html":
        *m = Html
    case "latex":
        *m = Latex
    default:
        err = errors.UnrecognizedMode(s, "table mode")
    }
    
    
    
    return
}
