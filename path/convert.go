package path

import (
    "github.com/bozso/gotoolbox/iter"
)

type IndexedFrom interface {
    GetFrom(ii int) From
}

type From interface {
    FromPath(p Pather) error
}

type Iterable interface {
    Iter() Iterator
}

type Iterator interface {
    Next() Pather
}

type Valids struct {
    iter.Indices
    paths []Valid
}

func newValids(paths []Valid) (v Valids) {
    return Valids{
        Indices: iter.Indices{Start: 0, Stop: len(paths),  Step: 1},
        paths: paths,
    }
}

func (v Valids) Iter() (it Iterator) {
    return &ValidIter{
        Iter: v.Indices.Iter(),
        paths: v.paths,
    }
}

type ValidIter struct {
    iter.Iter
    paths []Valid
}

func (v *ValidIter) Next() (p Pather) {
    p = v.paths[v.Iter.Current]
    
    if !v.Iter.Done() {
        p = nil
    }
    
    return 
}
