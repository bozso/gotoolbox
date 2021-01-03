package path

import (
    "os"
    "bufio"
    "errors"
    "io/ioutil"
    "encoding/json"
)

type File struct {
    Path
}

func (f File) IfExists() (opt *ValidFile, err error) {
    vf, err := f.ToValid()
    if err == nil {
        opt = &vf
    // file does not exist, optPath is nil, no error is raised
    } else if errors.Is(err, DoesNotExist) {
        err = nil
    }
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

func (p Path) ToValidFile() (vf ValidFile, err error) {
    v, err := p.ToValid()
    if err != nil {
        return
    }

    vf, err = v.ToFile()
    return
}

func (f File) ToValid() (vf ValidFile, err error) {
    v, err := f.Path.ToValid()
    if err != nil {
        return
    }

    vf, err = v.ToFile()
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
