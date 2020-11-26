package parsing

import (
    "fmt"
    "strconv"
)


type Int struct {
    inner int
}

func (i *Int) Set(s string) (err error) {
    i.inner, err = strconv.Atoi(s)
    if err != nil {
        err = NumberParseFail{
            Source: s,
            NumericalType: NumberInt,
            err: err,
        }
    }
    return
}

func (i Int) Get() (ii int) {
    return i.inner
}

type Float32 struct {
    inner float32
}

func (f *Float32) Set(s string) (err error) {
    fl, err := strconv.ParseFloat(s, 32)
    if err != nil {
        err = NumberParseFail{
            Source: s,
            NumericalType: NumberFloat32,
            err: err,
        }
    }

    f.inner = float32(fl)

    return
}

func (f Float32) Get() (fl float32) {
    return f.inner
}

type Float64 struct {
    inner float64
}

func (f *Float64) Set(s string) (err error) {
    f.inner, err = strconv.ParseFloat(s, 64)
    if err != nil {
        err = NumberParseFail{
            Source: s,
            NumericalType: NumberFloat64,
            err: err,
        }
    }

    return
}

func (f Float64) Get() (fl float64) {
    return f.inner
}

type NumberParseFail struct {
    NumericalType
    Source string
    err error
}

func (n NumberParseFail) Unwrap() (err error) {
    return n.err
}

func (n NumberParseFail) Error() (s string) {
    return fmt.Sprintf("failed to parse '%s' as a(n) %s", n.Source,
        n.NumericalType.String())

}
