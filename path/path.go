package path

import (
    "os"
    "strings"
    "path/filepath"
    pa "path"
)

type Pather interface {
    GetPath() string
}

type Baser interface {
    Base() string
}

type Path struct {
    s string
}

func New(p string) Path {
    return Path{p}
} 

func Joined(args ...string) Path {
    return Path{filepath.Join(args...)}
}

func (p Path) Join(args ...string) Path {
    ss := []string{p.s}
    
    return Joined(append(ss, args...)...)
}

func (p Path) Glob() (v []Valid, err error) {
    paths, err := filepath.Glob(p.GetPath())
    if err != nil {
        return
    }
    
    v = make([]Valid, 0)
    
    for _, p := range paths {
        pp, Err := New(p).ToValid()
        if Err != nil {
            err = Err
            return
        }
        
        v = append(v, pp)
    }
    
    return
}

func (p Path) GetPath() string {
    return p.s
}

func (p Path) String() string {
    return p.GetPath()
}

func (p Path) Abs() (pp Path, err error) {
    pp.s, err = filepath.Abs(p.String())
    
    if err != nil {
        err = p.Fail(OpCreateAbs, err)
    }
    
    return
}

func (p Path) Base() (pp Path) {
    pp.s = filepath.Base(p.String())
    return
}

func (p Path) Len() int {
    return len(p.String())
}

func (p Path) Ext() string {
    return pa.Ext(p.String())
}

func (p Path) NoExt() (pp Path) {
    s := p.String()
    
    return Path{strings.TrimSuffix(s, p.Ext())}
}

func (p Path) Create() (of *os.File, err error) {
    of, err = os.Create(p.String())
    
    if err != nil {
        err = p.Fail(OpCreate, err)
    }
    
    return 
}

func (p Path) Exist() (b bool, err error) {
    b, s := false, p.s
    _, err = os.Stat(s)

    if err == nil {
        b = true
        return
    }
    
    if os.IsNotExist(err) {
        err = nil
        return
    }
    
    err = p.Fail(OpExists, err)
    return
}

func (p Path) ToValid() (vp Valid, err error) {
    exist, err := p.Exist()
    if err != nil {
        return
    }
    
    if !exist {
        //error
    }
    
    vp.Path = p
    return
}


type ByModTime []Valid

func (a ByModTime) Len() int           { return len(a) }
func (a ByModTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a ByModTime) Less(i, j int) bool {
    return a[i].ModTime().Before(a[j].ModTime())
}
