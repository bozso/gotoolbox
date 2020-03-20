package path

import (
    "fmt"
    "os"
    "time"

    "github.com/bozso/gotoolbox/errors"
)

type Valid struct {
    Path
}

func (vp Valid) Stat() (fi os.FileInfo, err error) {
    fi, err = os.Lstat(vp.String())
    
    if err != nil {
        err = vp.Fail(OpStat, err)
    }
    
    return
}

func (vp Valid) ModTime() (t time.Time) {
    // what should happen in case of an error
    fi, _ := vp.Stat()
    
    return fi.ModTime()
}

func (vp Valid) IsDir() (b bool, err error) {
    b = false
    fi, err := vp.Stat()
    if err != nil {
        return
    }
    
    b = fi.IsDir()
    return
}

func (vp Valid) Open() (of *os.File, err error) {
    of, err = os.Open(vp.String())
    
    if err != nil {
        err = vp.Fail(OpOpen, err)
    }
    
    return 
}

func (vp Valid) Rename(target Pather) (dst Valid, err error) {
    s1, s2 := vp.String(), target.GetPath()
    
    if err = os.Rename(s1, s2); err != nil {
        err = errors.WrapFmt(err, "failed to move '%s' to '%s'", s1, s2)
        return
    }
    
    dst, err = New(s2).ToValid()
    return
}

func (vp Valid) Move(dir Dir) (dst Valid, err error) {
    _dst, err := dir.Join(vp.Base().String()).Abs()
    if err != nil {
        return
    }
    
    dst, err = vp.Rename(_dst)
    
    return
}

type validPathOperation int

const (
    OpOpen validPathOperation = iota
)

func (op validPathOperation) Fmt(p Path) (s string) {
    ps := p.GetPath()
    
    switch op {
    case OpOpen:
        s = fmt.Sprintf("failed to open path '%s'", ps)
    }
    return
}
