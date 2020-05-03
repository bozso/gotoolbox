package media

import (
    "github.com/bozso/gotoolbox/errors"
    "github.com/bozso/gotoolbox/path"
)

type Media struct {
    path.ValidFile
}

func New(vf path.ValidFile) (m Media) {
    m.ValidFile = vf
    return m
}

type Paths struct {
    content []path.ValidFile
}

func NewPaths(vf []path.ValidFile) (m Paths) {
    m.content = vf
    return
}

func (p Paths) Len() errors.Bound {
    return errors.Bound(len(p.content))
}

func (p Paths) Iter() []path.ValidFile {
    return p.content
}

func (p Paths) Index(ii int) (m Media, err error) {
    if err = p.Len().IsOutOfBounds(ii); err != nil {
        return
    }
    //fmt.Printf("%v\n", p.content[ii])
    
    return Media{p.content[ii]}, nil
}

func (p Paths) First() (m Media) {
    return Media{p.content[0]}
}

func (p Paths) Last() (m Media) {
    return Media{p.content[len(p.content) - 1]}
}
