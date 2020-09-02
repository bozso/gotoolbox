package doc

import (
    "github.com/bozso/gotoolbox/path"
)

type Loader struct {
    Status
    root path.Dir
}

func (_ Doc) NewFileLoader(args ...string) (l Loader) {
    l.root, l.Err  = path.Joined(args...).ToDir()
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

type Bytes struct {
    Status
    content []byte
}

type LoadEncoder struct {
    Loader
    encoder FileEncoder    
}

func (d Doc) Base64Encoder(args ...string) (le LoadEncoder) {
    return d.NewFileLoader(args...).WithEncoder(Base64)
}

func (l Loader) WithEncoder(f FileEncoder) (le LoadEncoder) {
    le.Loader, le.encoder = l, NoOpEncoder()
    return
}

func (le LoadEncoder) Encode(args ...string) (str StringResult) {
    if str.From(le) {
        return
    }
    
    p := le.Retreive(args...)
    if p.Err != nil {
        return
    }
    
    f, err := p.Valid.ToFile()
    if err != nil {
        str.Err = err
        return
    }
    
    s, err := le.encoder.EncodeFile(f)
    if err != nil {
        p.Err = err
        return
    } else {
        str.s = s
    }

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

type StringResult struct {
    Status
    s string
}

func OkString(s string) (sr StringResult) {
    sr.s = s
    return
}

func (str StringResult) String() (s string) {
    return str.s
}
