package assets

import (
	"strings"

	"git.sr.ht/~istvan_bozso/shutil/fs"
	"github.com/bozso/gotoolbox/enum"
)

var osFS = &fs.OsFS{}

type PathType int

const (
	Absolute PathType = iota
	Relative
)

var pathEnum = enum.NewStringSet("absolute", "relative").EnumType("PathType")

func (p *PathType) Set(s string) (err error) {
	switch strings.ToLower(s) {
	case "absolute":
		*p = Absolute
	case "relative":
		*p = Relative
	default:
		err = pathEnum.UnknownElement(s)
	}

	return
}

type Path struct {
	Path string   `json:"path"`
	Type PathType `json:"type"`
}

func (p Path) Open(fsys fs.FS) (f fs.File, err error) {
	if p.Type == Absolute {
		fsys = osFS
	}

	return fsys.Open(p.Path)
}

func (p Path) OpenMut(fsys fs.MutFS) (f fs.MutFile, err error) {
	if p.Type == Absolute {
		fsys = osFS
	}

	return fsys.OpenMut(p.Path)
}
