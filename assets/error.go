package assets

import (
	"fmt"
	"net/url"
)

type Operation string

func (op Operation) Error(err error) (e Error) {
	return Error{
		Operation: op,
		err:       err,
	}
}

type Error struct {
	Operation
	err error
}

func (e Error) Error() (s string) {
	return fmt.Sprintf("while carrying out operation: %s", string(e.Operation))
}

func (e Error) Unwrap() (err error) {
	return e.err
}

type DownloadError struct {
	*url.URL
	err error
}

func (d DownloadError) Error() (s string) {
	return fmt.Sprintf("downloading asset from URL: '%s' failed", d.URL)
}

func (d DownloadError) Unwrap() (err error) {
	return d.err
}
