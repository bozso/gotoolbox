package parsing

import (
    "fmt"
)

type Getter interface {
    Get(string) string
}

type Value interface {
    Set(string) error
}

type GetterParser struct {
    getter Getter
}




func (g GetterParser) Get(key string) (s string, b bool) {
    s = g.getter.Get(key)

    if len(s) != 0 {
        b = true
    }
    return
}

func (g GetterParser) GetVal(key string, val Value) (err error) {
    s, err := g.MustGet(key)
    if err != nil {
        return
    }

    err = val.Set(s)
    return
}

func (g GetterParser) GetInt(key string) (i int, err error) {
    var ii Int
    err = g.GetVal(key, &ii)
    if err != nil {
        return
    }

    i = ii.Get()
    return
}

func (g GetterParser) GetFloat32(key string) (f float32, err error) {
    var fl Float32
    err = g.GetVal(key, &fl)
    if err != nil {
        return
    }

    f = fl.Get()
    return
}

func (g GetterParser) GetFloat64(key string) (f float64, err error) {
    var fl Float64
    err = g.GetVal(key, &fl)
    if err != nil {
        return
    }

    f = fl.Get()
    return
}

func (g GetterParser) GetFloat(key string) (f float64, err error) {
    return g.GetFloat64(key)
}

func (g GetterParser) MustGet(key string) (s string, err error) {
    s, b := g.Get(key)

    if !b {
        err = KeyNotFound{key}
    }

    return
}

type KeyNotFound struct {
    Key string
}

func (k KeyNotFound) Error() (s string) {
    return fmt.Sprintf("key '%s' not found", k.Key)
}
