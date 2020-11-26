package parsing

import (
)

type NumericalType int

const (
    NumberInt NumericalType = iota
    NumberFloat32
    NumberFloat64
)

func (n NumericalType) String() (s string) {
    switch n {
    case NumberInt:
        s = "integer"
    case NumberFloat32:
        s = "float32"
    case NumberFloat64:
        s = "float64"
    default:
        s = "unknown"
    }
    return
}
