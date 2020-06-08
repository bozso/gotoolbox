package path

import (
    "os"
    "bufio"
    "io/ioutil"
    "encoding/json"

    "github.com/bozso/gotoolbox/errors"
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
    v, err := f.Path.ToValid()
    if err != nil {
        return
    }
    
    vf, err = v.ToFile()
    return
}

func (vf *ValidFile) Set(s string) (err error) {
    const name errors.NotEmpty = "valid file path"

    if err = name.Check(s); err != nil {
        return
    }
    
    *vf, err = NewFile(s).ToValid()
    return
}

func (vf *ValidFile) UnmarshalJSON(b []byte) (err error) {
    err = vf.Set(trim(b))
    return
}

func (vf ValidFile) ToFile() (f File) {
    f.Path = vf.Path
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

func (vf ValidFile) LoadJson(val interface{}) (err error) {
    b, err := vf.ReadAll()
    if err != nil {
        return
    }
    
    err = json.Unmarshal(b, val)
    if err != nil {
        return
    }
    
    v, ok := val.(Validator)
    if !ok {
        return
    }
    
    err = v.Validate()
    return
}

type Validator interface {
    Validate() error
}
