package path

import (
    "fmt"
    "io/ioutil"
)

type File struct {
    Path
}

func (p Path) ToFile() (f File, err error) {
    fi, err := p.Stat()
    
    if fi.IsDir() {
        err = fmt.Errorf("path '%s' is not a file, but a directory", p)
        return
    }
    
    f.Path = p
    
    return
}

func (f File) ReadAll() (b []byte, err error) {
    file, err := f.Open()
    if err != nil {
        return
    }
    defer file.Close()
    
    b, err = ioutil.ReadAll(file)
    return
}

func (f *File) Set(s string) (err error) {
    *f, err = New(s).ToFile()
    return
}

