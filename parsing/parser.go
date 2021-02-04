package parsing

import (
    "fmt"
)

type Getter interface {
    GetParam(string) string
}

type Value interface {
    Set(string) error
}


func Get(g Getter, key string) (s string, b bool) {
    s = g.GetParam(key)

    if len(s) != 0 {
        b = true
    }
    return
}

func GetVal(g Getter, key string, val Value) (err error) {
    s, err := MustGet(g, key)
    if err != nil {
        return
    }

    err = val.Set(s)
    return
}

func GetInt(g Getter, key string) (i int, err error) {
    var ii Int
    err = GetVal(g, key, &ii)
    if err != nil {
        return
    }

    i = ii.Get()
    return
}

func GetFloat32(g Getter, key string) (f float32, err error) {
    var fl Float32
    err = GetVal(g, key, &fl)
    if err != nil {
        return
    }

    f = fl.Get()
    return
}

func GetFloat64(g Getter, key string) (f float64, err error) {
    var fl Float64
    err = GetVal(g, key, &fl)
    if err != nil {
        return
    }

    f = fl.Get()
    return
}

func GetFloat(g Getter, key string) (f float64, err error) {
    return GetFloat64(g, key)
}

func MustGet(g Getter, key string) (s string, err error) {
    s, b := Get(g, key)

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
