package path

import (
    "os"
    "fmt"
    "io/ioutil"
)

type Dir struct {
    Valid
}

func (d Dir) Remove() (err error) {
    err = os.RemoveAll(d.GetPath())
    return
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

func (p Path) Mkdir() (d Dir, err error) {
    if err = os.MkdirAll(p.GetPath(), os.ModePerm); err != nil {
        err = d.Fail(DirCreate, err)
    }
    
    d.Path = p
    return
}

func (d Dir) Chdir() (err error) {
    return os.Chdir(d.String())
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

func TempDirIn(dir, prefix string) (d Dir, err error) {
    dirPath, err := ioutil.TempDir(dir, prefix)
    if err != nil {
        return
    }
    
    d, err = New(dirPath).ToDir()
    return
}

func TempDir(prefix string) (d Dir, err error) {
    d, err = TempDirIn("", prefix)
    return
}

type EmptyDir struct {
    Dir
}

func (e *EmptyDir) SetPath(p Path) {
    e.Dir.Path = p
}

func (_ EmptyDir) Create(p Path) (err error) {
    _, err = p.Mkdir()
    return
}

func (e *EmptyDir) Set(s string) (err error) {
    err = CreateIf(e, s)
    return
}

func (e *EmptyDir) UnmarshalJSON(b []byte) (err error) {
    err = e.Set(trim(b))
    return
}
