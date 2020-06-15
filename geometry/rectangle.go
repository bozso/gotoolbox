package geometry

import (
    "io"
    "fmt"
)

type Point2D struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type LeftRight struct {
    Left  Point `json:"left"`
    Right Point `json:"right"`
}

type Rectangle struct {
    Upper LeftRight `json:"upper"`
    Lower LeftRight `json:"lower"`
}

func (r Rectangle) InitFormat(wr io.Writer) (n int, err error) {
    // TODO: rework tpl string
    const tpl = "{0:>12s}{1:>12s}{2:>12s}{3:>12s}{4:>12s}{5:>12s}{6:>12s}{7:>12s}"

    s := fmt.Sprintf(tpl,
        r.Lower.Left.X, r.Lower.Left.Y,
        r.Upper.Left.X, r.Upper.Left.Y,
        r.Lower.Right.X, r.Lower.Right.Y,
        r.Upper.Right.X, r.Upper.Right.Y)
    
    return wr.Write([]byte(s))
}
