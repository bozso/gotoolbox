package path

import (
    "fmt"
    "os"
    "bufio"
    "io/ioutil"
)

type File struct {
    Path
}

func NewFile(s string) (f File) {
    f.Path = New(s)
    return
}

func (f *File) Set(s string) (err error) {
    *f, err = NewFile(s), nil
    return
}

func (p Path) ToFile() (f File) {
    f.Path = p    
    return
}

func (p Path) Remove() (err error) {
    err = os.Remove(p.GetPath())
    return
}

type ValidFile struct {
    Valid
}

func (f File) ToValid() (vf ValidFile, err error) {
    vf.Valid, err = f.Path.ToValid()
    if err != nil {
        return
    }
    
    isDir, err := vf.IsDir()
    if err != nil {
        return
    }
    
    if isDir {
        err = fmt.Errorf("path '%s' is not a file, but a directory", f)
    }

    return
}

func (vf *ValidFile) Set(s string) (err error) {
    *vf, err = NewFile(s).ToValid()
    return
}

func (vf ValidFile) Move(d Dir) (vfm ValidFile, err error) {
    vfm.Valid, err = vf.Valid.Move(d)
    return
}

type Scanner struct {
    *bufio.Scanner
    *os.File
}

func (vf ValidFile) Scanner() (s Scanner, err error) {
    s.File, err = vf.Open()
    if err != nil {
        return
    }
    
    s.Scanner = bufio.NewScanner(s.File)
    return
}

func (vf ValidFile) ReadAll() (b []byte, err error) {
    file, err := vf.Open()
    if err != nil {
        return
    }
    defer file.Close()
    
    b, err = ioutil.ReadAll(file)
    return
}
