package path

import (
	"fmt"
)

type Extension string

func (e Extension) Check(v ValidFile) (b bool) {
	return v.Extension() == e
}

func (e Extension) MustHave(v ValidFile) (err error) {
	if ext := v.Extension(); ext != e {
		err = WrongExtension{
			Expected: e,
			Got:      ext,
			File:     v,
		}
	}
	return
}

type MustHaveExtension Extension

/*
WrongExtension represents an error that occurs when a valid file does not
have the proper extension.
*/
type WrongExtension struct {
	// The expected extension.
	Expected Extension
	// The extension we got.
	Got Extension
	// The file with the wrong extension.
	File ValidFile
}

func (e WrongExtension) Error() (s string) {
	return fmt.Sprintf(
		"expected extension '%s' got '%s' for path '%s'", e.Expected,
		e.Got, e.File)
}

func (e Extension) Files(num int) (w WithExtension) {
	return WithExtension{
		files:     make([]ValidFile, num),
		extension: e,
	}
}

/*
WithExtension represents a slice of files that have the proper extension.
Implements the IndexedFrom interface. An easy way to generate a slice
of valid file paths with the right extension.
*/
type WithExtension struct {
	files     []ValidFile
	extension Extension
}

// Files returns the valid file paths that have the right extension.
func (w WithExtension) Files() (files []ValidFile) {
	return w.files
}

func (w WithExtension) GetFrom(ii int) (f From) {
	return &withExtension{
		file:      &w.files[ii],
		extension: w.extension,
	}
}

type withExtension struct {
	file      *ValidFile
	extension Extension
}

func (w *withExtension) FromPath(p Pather) (err error) {
	*w.file, err = p.AsPath().ToValidFile()
	if err != nil {
		return
	}

	err = w.extension.MustHave(*w.file)
	return
}
