package extract

import (
	"io/fs"
	"path/filepath"
)

type Extractor interface {
	Extract(input, outdir string) error
}

type Creator interface {
	CreateFs(path string) (fs.FS, error)
}

type Type int

const (
	Zip Type = iota
	Tar
	Gzip
)

func Identify(path string) (t Type, b bool) {
	b = true

	switch filepath.Ext(path) {
	case "tar":
		t = Tar
	case "gz":
		t = Gzip
	case "zip":
		t = Zip
	case "":
		b = false
	}

	return
}

func MustIdentify(path string) (t Type, err error) {
	return
}
