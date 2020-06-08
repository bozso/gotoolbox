package media

import (
    "github.com/bozso/gotoolbox/errors"
    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/doc/result"
)

type Paths struct {
    result.Status
    content []path.ValidFile
}

func NewPaths(vf []path.ValidFile) (p Paths) {
    p.content = vf
    return
}

func (p Paths) Len() errors.Bound {
    return errors.Bound(len(p.content))
}

func (p Paths) Iter() []path.ValidFile {
    return p.content
}

func (p Paths) Index(ii int) (r Result) {
    if !p.IsValid() {
        r.Status = p.Status
        return
    }

    if err := p.Len().IsOutOfBounds(ii); err != nil {
        r.Status = result.Error(err)
        return 
    }

    return New(p.content[ii])
}

func (p Paths) First() (r Result) {
    if !p.IsValid() {
        r.Status = p.Status
        return
    }

    return New(p.content[0])
}

func (p Paths) Last() (r Result) {
    if !p.IsValid() {
        r.Status = p.Status
        return
    }
    return New(p.content[len(p.content) - 1])
}
