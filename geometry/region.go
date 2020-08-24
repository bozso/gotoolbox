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
    Min Point2D `json:"min"`
    Max Point2D `json:"max"`
}

func NewRegion(xmin, xmax, ymin, ymax float64) (r Region) {
    return Region{
        Min: Point2D{
            X: xmin,
            Y: ymin,
        },
        Max: Point2D{
            X: xmax,
            Y: ymax,
        },
    }
}

func (r Region) Contains(p Point2D) (b bool) {
    return (p.X < r.Max.X && p.X > r.Min.X &&
            p.Y < r.Max.Y && p.Y > r.Min.Y)
}

func (r Region) Upper() (lr LeftRight2D) {
    lr.Left.X, lr.Left.Y = r.Min.X, r.Max.Y
    lr.Right.X, lr.Right.Y = r.Min.X, r.Max.Y
    return
}

func (r Region) Lower() (lr LeftRight2D) {
    lr.Left.X, lr.Left.Y = r.Min.X, r.Min.Y
    lr.Right.X, lr.Right.Y = r.Max.X, r.Min.Y
    return
}
