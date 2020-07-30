package server

import (
    "fmt"
)

type Version string

type VersionMismatch struct {
    expected Version
    got string
}

func (v VersionMismatch) Error() (s string) {
    return fmt.Sprintf("Expected version string to be '%s', got '%s'",
        v.expected, v.got)
}

func (v Version) Validate(s string) (err error) {
    if s != string(v) {
        err = VersionMismatch{expected: v, got: s}
    }
    return
}
