package geometry

import (
)

type Point2D struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type LeftRight2D struct {
    Left  Point2D `json:"left"`
    Right Point2D `json:"right"`
}

type Rectangle struct {
    X MinMaxFloat `json:"x"`
    Y MinMaxFloat `json:"y"`
}

func (r Rectangle) Contains(p Point2D) (b bool) {
    return p.InRectangle(r)
}

func (p Point2D) InRectangle(r Rectangle) (b bool) {
    return (p.X < r.X.Max && p.X > r.X.Min &&
            p.Y < r.Y.Max && p.Y > r.Y.Min)
}
