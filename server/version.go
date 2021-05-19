package server

import (
	"fmt"
)

/*
An alias for a string representing a version.
*/
type Version string

type VersionMismatch struct {
	expected Version
	got      string
}

// Implements the error interface.
func (v VersionMismatch) Error() (s string) {
	return fmt.Sprintf("Expected version string to be '%s', got '%s'",
		v.expected, v.got)
}

/*
Compare the two version strings.
*/
func (v Version) Validate(s string) (err error) {
	if s != string(v) {
		err = VersionMismatch{expected: v, got: s}
	}
	return
}
