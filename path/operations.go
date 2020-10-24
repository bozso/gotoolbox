package path

import (
    "github.com/bozso/gotoolbox/errors"
)

type Glob func() ([]Valid, error)
type Filter func(Pather) (bool, error)
type EnumFilter func(int, Pather) (bool, error)
type Transform func(Pather) (Pather, error)

type Filters []Filter
type Transforms []Transform

func SelectAll(_ Pather) bool {
    return true
}

type Filterer interface {
    Filter(Pather) (bool, error) 
}

type SelectDir struct {}

func (_ SelectDir) Filter(p Pather) (b bool, err error) {
    v, err := p.AsPath().ToValid()
    if err != nil {
        return
    }
    
    b, err = v.IsDir()
    return
    
}

type SelectFile struct {}

func (_ SelectFile) Filter(p Pather) (b bool, err error) {
    v, err := p.AsPath().ToValid()
    if err != nil {
        return
    }
    
    b, err = v.IsDir()
    b = !b
    return
    
}

type selectAll struct{}

func (_ selectAll) Filter(_ Pather) (b bool, err error) {
    return true, nil
}

var selectors = map[string]Filterer{
    "dirs": SelectDir{},
    "files": SelectFile{},
    "all": selectAll{},
}

type FiltererPayload struct {
    Filterer
}

func (f *FiltererPayload) UnmarshalJSON(b []byte) (err error) {
    s := trim(b)
    var ok bool
    f.Filterer, ok = selectors[s]
    
    if !ok {
        err = errors.KeyNotFound(s)
    }
    return
}
