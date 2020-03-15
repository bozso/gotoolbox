package path

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
    pa "path"

    "github.com/bozso/gotoolbox/errors"
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

func (p Path) GetPath() string {
    return p.s
}

func (p Path) String() string {
    return p.GetPath()
}

func (p Path) Abs() (pp Path, err error) {
    pp.s, err = filepath.Abs(p.String())
    
    if err != nil {
        err = p.Fail(CreateAbsPath, err)
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
        err = p.Fail(Create, err)
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
    
    err = errors.WrapFmt(err,
        "failed to check wether file '%s' exists", s)
    return
}

func (p Path) Fail(op PathOperation, err error) error {
    return PathError{p.s, op, err}
}

func (p Path) ToValid() (vp validPath, err error) {
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

type validPath struct {
    Path
}

func (vp validPath) Stat() (fi os.FileInfo, err error) {
    fi, err = os.Lstat(vp.String())
    
    if err != nil {
        err = vp.Fail(Stat, err)
    }
    
    return
}

func (vp validPath) IsDir() (b bool, err error) {
    b = false
    fi, err := vp.Stat()
    if err != nil {
        return
    }
    
    b = fi.IsDir()
    return
}

func (vp validPath) Open() (of *os.File, err error) {
    of, err = os.Open(vp.String())
    
    if err != nil {
        err = vp.Fail(Open, err)
    }
    
    return 
}

func (vp validPath) Move(dir Dir) (dst validPath, err error) {
    _dst, err := dir.Join(vp.Base().String()).Abs()
    if err != nil {
        return
    }
    
    s1, s2 := vp.String(), dst.String()
    
    if err = os.Rename(s1, s2); err != nil {
        errors.WrapFmt(err, "failed to move '%s' to '%s'", s1, s2)
        return
    }
    
    dst, err = _dst.ToValid()
    
    return
}

type PathOperation int

const (
    Stat PathOperation = iota
    Open
    Create
    CheckIfExist
    CreateAbsPath
)

func (op PathOperation) String() (s string) {
    switch op {
    case Stat:
        s = "retreive information"
    case Open:
        s = "open"
    case Create:
        s = "create"
    case CheckIfExist:
        s = "check if exists"
    case CreateAbsPath:
        s = "create absolute path"
    }
    
    return
}

type PathError struct {
    s string
    op fmt.Stringer
    err error
}

func (e PathError) Error() string {
    return fmt.Sprintf(
        "failed to carry out operation '%s' on path '%s'",
         e.op.String(), e.s)
}

func (e PathError) Unwrap() error {
    return e.err
}
