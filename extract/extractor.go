package extract

import (
	"io/fs"
)

type Extractor interface {
	NewFs(path string) (fs.FS, error)
}

type Type int

const (
	Zip Type = iota
	Tar
)

func Identify(path string) (t Type, b bool) {
	return
}

func MustIdentify(path string) (t Type, err error) {
	return
}
