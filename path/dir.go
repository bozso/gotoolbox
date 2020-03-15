package path

import (
    "os"
    "fmt"
)

type Dir struct {
    validPath
}

func (p Path) ToDir() (d Dir, err error) {
    d.validPath, err = p.ToValid()
    if err != nil {
        return
    }
    
    isDir, err := d.IsDir()
    if err != nil {
        return
    }
    
    if !isDir {
        err = fmt.Errorf("path '%s' is not a directory, but a file", p)
    }
    
    return
}

func (d *Dir) Set(s string) (err error) {
    *d, err = New(s).ToDir()
    return
}


func (d Dir) Mkdir() (err error) {
    s := d.s
    if err = os.MkdirAll(s, os.ModePerm); err != nil {
        err = d.Fail(create, err)
    }
    
    return
}

type DirOperation int

const (
    create DirOperation = iota
)

func (op DirOperation) String() (s string) {
    switch op {
    case create:
        s = "create directory"
    }
    return s
}

func (d Dir) Fail(op DirOperation, err error) error {
    return PathError{d.s, op, err}
}
