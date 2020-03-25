package path

import (
    "os"
    "fmt"
    "io/ioutil"
)

type Dir struct {
    Valid
}

func (p Path) ToDir() (d Dir, err error) {
    d.Valid, err = p.ToValid()
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

func (d *Dir) UnmarshalJSON(b []byte) (err error) {
    err = d.Set(trim(b))
    return
}

func (d Dir) Mkdir() (err error) {
    s := d.s
    if err = os.MkdirAll(s, os.ModePerm); err != nil {
        err = d.Fail(DirCreate, err)
    }
    
    return
}

func (d Dir) ReadDir() (fi []os.FileInfo, err error) {
    fi, err = ioutil.ReadDir(d.GetPath())
    return
}

type dirOperation int

const (
    DirCreate dirOperation = iota
)

func (op dirOperation) Fmt(p Path) (s string) {
    ps := p.GetPath()
    
    switch op {
    case DirCreate:
        s = fmt.Sprintf("failed to create directory '%s'", ps)
    }    
    
    return 
}
