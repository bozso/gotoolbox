package repository

import (
    "github.com/bozso/gotoolbox/path"
)

type Repository struct {
    Directory path.Dir `json:"directory"`
}

func (r *Repository) Set(s string) (err error) {
    return r.FromPath(path.New(s))
}

func (r *Repository) UnmarshalJSON(b []byte) (err error) {
    return r.Directory.UnmarshalJSON(b)
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

type Repositories []Repository
