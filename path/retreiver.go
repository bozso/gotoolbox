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

func (r Retreiver) Format(args ...interface{}) (p File) {
    return r.searchDir.Join(fmt.Sprintf(r.pattern, args...)).ToFile()
}

func (r Retreiver) Index(idx int) (vf ValidFile, err error) {
    vf, err = r.Format(idx).ToValid()
    return
}

func (r Retreiver) Key(key string) (vf ValidFile, err error) {
    vf, err = r.Format(key).ToValid()
    return
}

func (r Retreiver) Get() (vf ValidFile, err error) {
    vf, err = r.Format().ToValid()
    return
}
