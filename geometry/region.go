package geometry

import (
)

type Axis int

const (
    X Axis = iota
    Y
    AxisNum
)

type MinMax int

const (
    Min MinMax = iota
    Max
    MinMaxNum
)

type MinMaxF64 [MinMaxNum]float64

type MinMaxFloat struct {
    Min float64 `json:"min"`
    Max float64 `json:"max"`
}

type Region struct {
    X MinMaxFloat `json:"x"`
    Y MinMaxFloat `json:"y"`
}

func NewRegion(xmin, xmax, ymin, ymax float64) (r Region) {
    return Region{
        X: MinMaxFloat{
            Min: xmin,
            Max: xmax,
        },
        Y: MinMaxFloat{
            Min: ymin,
            Max: ymax,
        },
    }
}
