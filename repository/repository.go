package repository

import (
    "github.com/bozso/gotoolbox/path"
)

type Repository struct {
    Directory path.Dir
}

func New(dir path.Dir) (r Repository) {
    return Repository {
        Directory: dir,
    }
}

func (r *Repository) FromPath(p path.Pather) (err error) {
    d, err := p.AsPath().ToDir()
    
    *r = New(d)
    return
}
