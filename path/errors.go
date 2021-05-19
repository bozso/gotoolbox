package path

import (
	"fmt"
<<<<<<< HEAD

	"github.com/bozso/gotoolbox/errors"
=======
>>>>>>> ee017e5a8603787d06b1d4a7f1f8a150c8268130
)

var Error pathError

func (p Path) Fail(efmt ErrorFmt, err error) (Err error) {
	return pathError{p, efmt, err}
}

type ErrorFmt interface {
	Fmt(Path) string
}

type pathError struct {
	p    Path
	efmt ErrorFmt
	err  error
}

<<<<<<< HEAD
func (e pathError) Unwrap() error {
=======
func (e pathError) Unwrap() (err error) {
>>>>>>> ee017e5a8603787d06b1d4a7f1f8a150c8268130
	return e.err
}

func (e pathError) Error() string {
<<<<<<< HEAD
	return errors.WrapFmt(e.Unwrap(), "%s", e.efmt.Fmt(e.p)).Error()
=======
	return e.efmt.Fmt(e.p)
>>>>>>> ee017e5a8603787d06b1d4a7f1f8a150c8268130
}

type operation int

const (
	OpCreate operation = iota
	OpCreateAbs
	OpStat
	OpExists
)

func (op operation) Fmt(P Path) (s string) {
	p := P.GetPath()

	switch op {
	case OpCreate:
		s = fmt.Sprintf("failed to create path '%s'", p)
	case OpCreateAbs:
		s = fmt.Sprintf("failed to create absolute path '%s'", p)
	case OpStat:
		s = "failed to retreive information:"
	case OpExists:
		s = fmt.Sprintf("failed to check if path '%s' exists", p)
		// add default case?
	}
	return
}
