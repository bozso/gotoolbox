package cli

import (
    "fmt"
)

type ColorStr string

const (
    Info    ColorStr = "\033[1;34m"
    Notice  ColorStr = "\033[1;36m"
    Warning ColorStr = "\033[1;33m"
    Error   ColorStr = "\033[1;31m"
    Debug   ColorStr = "\033[0;36m"
    Bold    ColorStr = "\033[1;0m"
    End     ColorStr = "\033[0m"
    
)

func (c ColorStr) Wrap(s string) (out string) {
    out = fmt.Sprintf("%s%s%s", string(c), s, string(End))
    return
}

func (c ColorStr) Fmt(msg string, args ...interface{}) (s string) {
    s = c.Wrap(fmt.Sprintf(msg, args...))
    return
}
