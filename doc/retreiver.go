package doc

import (
    "github.com/bozso/gotoolbox/path"
)

type Loader struct {
    Status
    root path.Dir
}

func (_ Doc) NewFileLoader(args ...string) (l Loader) {
    l.root, l.Err = path.Joined(args...).ToDir()
    return
}

func (l Loader) Join(args ...string) (ll Loader) {
    if ll.From(l) {
        return
    }
    
    ll.root, ll.Err = l.root.Join(args...).ToDir()
    return
}

func (l Loader) Retreive(args ...string) (p Path) {
    if p.From(l) {
        return
    }
    
    p.TryFrom(l.root.Join(args...))
    return
}

type Path struct {
    Status
    path.Valid
}

func (p *Path) TryFrom(systemPath path.Path) {
    p.Valid, p.Err = systemPath.ToValid()
    return
}
