package server

import (
	"fmt"
)

type ConvertFailLabel string

func (c ConvertFailLabel) New() (cf ConvertFail) {
	cf.name = string(c)
	return
}

/*
Create a ConvertFail error.
*/
func ConversionFailure(str fmt.Stringer) (cf ConvertFail) {
	cf.name = str.String()
	return
}

/*
Represents a mismatch between an expected database type entry and the
actual database entry type.
*/
type ConvertFail struct {
	name string
}

// Implement error interface.
func (c ConvertFail) Error() (s string) {
	return fmt.Sprintf("failed to convert database entity to %s",
		c.name)
}
