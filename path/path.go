package path

import (
    "fmt"
    "os"
    "strings"
    "bytes"
    "path/filepath"
    "errors"
    pa "path"

    gerrors "github.com/bozso/gotoolbox/errors"
)

type Pather interface {
    AsPath() Path
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

func (p *Path) Set(s string) (err error) {
    *p = New(s)
    return
}

/*
 * Make all structs that embedd Path trivially convertable back to Path.
 */
func (p Path) AsPath() (rp Path) {
    return p
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

func (p Path) AddExt(ext string) (pp Path) {
    return New(fmt.Sprintf("%s.%s", p.GetPath(), ext))
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

func (p Path) Dir() (pp Path) {
    return New(filepath.Dir(p.GetPath()))
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

type NotExists string

func (e NotExists) Error() (s string) {
    return fmt.Sprintf("path '%s' does not exist", string(e))
}

const DoesNotExist NotExists = "" 

func (p Path) MustExist() (err error) {
    b, err := p.Exist()
    if err != nil {
        return
    }
    
    if !b {
        err = NotExists(p.String())
    }
    
    return
}

func (p Path) IfExists() (optPath *Valid, err error) {
    v, err := p.ToValid()
    
    if err == nil {
        optPath = &v
    // file does not exist, optPath is nil, no error is raised
    } else if errors.Is(err, DoesNotExist) {
        err = nil
    }
    
    return
}

func (p Path) ToValid() (vp Valid, err error) {
    if err = p.MustExist(); err != nil {
        return
    }
    
    vp.Path = p
    return
}

func (p Path) MarshalJSON() (b []byte, err error) {
    return []byte(fmt.Sprintf("\"%s\"", p.GetPath())), nil
}

func trim(b []byte) (s string) {
    return string(bytes.Trim(b, "\""))
}

func (p *Path) UnmarshalJSON(b []byte) (err error) {
    p.s = trim(b)
    return nil
}

type ByModTime []Valid

func (b ByModTime) Len() int           { return len(b) }
func (b ByModTime) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }


func (b ByModTime) Less(i, j int) bool {
    const check gerrors.Asserter = "failed to retreive modification time" 
    
    v1 := b[i]
    t1, err := v1.ModTime()
    check.Checkf(err, "for path '%s'", v1)

    v2 := b[j]
    t2, err := v2.ModTime()
    check.Checkf(err, "for path '%s'", v2)
    
    
    return t1.Before(t2)
}

type Creatable interface {
    SetPath(Path)
    Create(Path) error
}

func CreateIf(create Creatable, s string) (err error) {
    p := New(s)

    exists, err := p.Exist()
    if err != nil {
        return
    }
    
    if !exists {
        if err = create.Create(p); err != nil {
            return
        }
    }
    
    create.SetPath(p)
    return
} 
