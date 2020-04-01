package template

import (
    "github.com/bozso/gotoolbox/errors"
    "github.com/bozso/gotoolbox/path"
)

type Media struct {
    path.ValidFile
}

func NewMedia(vf path.ValidFile) (m Media) {
    m.ValidFile = vf
    return m
}

type MediaPaths struct {
    content []path.ValidFile
}

func NewPaths(vf []path.ValidFile) (m MediaPaths) {
    m.content = vf
    return
}

func (p MediaPaths) Len() errors.Bound {
    return errors.Bound(len(p.content))
}

func (p MediaPaths) Iter() []path.ValidFile {
    return p.content
}

func (p MediaPaths) Index(ii int) (m Media, err error) {
    if err = p.Len().IsOutOfBounds(ii); err != nil {
        return
    }
    //fmt.Printf("%v\n", p.content[ii])
    
    return Media{p.content[ii]}, nil
}

func (p MediaPaths) Last() (m Media) {
    return Media{p.content[len(p.content) - 1]}
}
