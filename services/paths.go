package services

import (
	"github.com/bozso/gotoolbox/hash"
	"github.com/bozso/gotoolbox/path"
)

type Empty struct{}

type Path struct {
	Path path.Path `json:"path"`
}

func (p Path) Hash(h hash.Hash) {
	p.Path.Hash(h)
}

type Paths struct {
	Paths []path.Path `json:"paths"`
}

func (p Paths) Hash(h hash.Hash) {
	for ii, _ := range p.Paths {
		p.Paths[ii].Hash(h)
	}
}
