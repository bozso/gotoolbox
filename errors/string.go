package errors

import (
	"fmt"
)

type NotEmpty string

func (n NotEmpty) Check(variable string) (err error) {
	if len(variable) == 0 {
		return EmptyStringError{string(n), nil}
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

const defaultCap = 4

type Strings struct {
	elems []string
}

func NewStringsWithCap(cap int) (s Strings) {
	s.elems = make([]string, 0, cap)
	return
}

func NewStrings() (s Strings) {
	return NewStringsWithCap(defaultCap)
}

func (s *Strings) Add(str string) {
	s.elems = append(s.elems, str)
}

func (s Strings) String() (str string) {
	return fmt.Sprintf("%v", s.elems)
}

type Name string

func (n Name) NewAvailable() (a Available) {
	a.Name, a.available = n, NewStrings()
	return
}

type Available struct {
	Name
	available Strings
}

func (a *Available) Add(s string) {
	a.available.Add(s)
}

func (a Available) NotFound(requested string) (nf NotFound) {
	nf.Available, nf.requested = a, requested
	return
}

type NotFound struct {
	requested string
	Available
}

func (e NotFound) Error() (s string) {
	return fmt.Sprintf("%s '%s' not found. Available %s(s): %s",
		e.Name, e.requested, e.Name, e.available)
}
