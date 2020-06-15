package geometry

import (
    "io"
    "fmt"
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

type Region struct {
    X, Y MinMaxF64
}

func NewRegion(xmin, xmax, ymin, ymax float64) (r Region) {
    return Region{
        X: MinMaxF64{
            Min: xmin,
            Max: xmax,
        },
        Y: MinMaxF64{
            Min: ymin,
            Max: ymax,
        },
    }
}

func (r Region) InitFormat(wr io.Writer) (n int, err error) {
    // TODO: rework tpl string
    const tpl = "{0:>22s}{1:>14s}{2:>12s}{3:>14s}"

    s := fmt.Sprintf(tpl,
        r.X[Min], r.X[Max],
        r.Y[Min], r.Y[Max])

    return wr.Write([]byte(s))
}

func (r Region) DirName() (s string) {
    return fmt.Sprintf("x_%d_%d__y_%d_%d",
        int(r.X[Min]), int(r.X[Max]),
        int(r.Y[Min]), int(r.Y[Max]))
}
