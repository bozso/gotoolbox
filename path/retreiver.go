package path

import (
    "fmt"
)

type Retreiver struct {
    searchDir Dir
    pattern string
}

func NewRetreiver(dir Dir, pattern string) (r Retreiver) {
    return Retreiver{
        searchDir: dir,
        pattern: pattern,
    }
}

func (r Retreiver) Format(args ...interface{}) (p File, err error) {
    return r.searchDir.Join(fmt.Sprintf(r.pattern, args...)).ToFile()
}

func (r Retreiver) Index(idx int) (vf ValidFile, err error) {
    f, err := r.Format(idx)
    if err != nil {
        return
    }

    return f.ToValid()
}

func (r Retreiver) Key(key string) (vf ValidFile, err error) {
    f, err := r.Format(key)
    if err != nil {
        return
    }

    return f.ToValid()
}

func (r Retreiver) Get() (vf ValidFile, err error) {
    f, err := r.Format()
    if err != nil {
        return
    }

    return f.ToValid()
}
