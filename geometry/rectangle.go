package geometry

import (
)

type Point2D struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type LeftRight struct {
    Left  Point2D `json:"left"`
    Right Point2D `json:"right"`
}

type Rectangle struct {
    Upper LeftRight `json:"upper"`
    Lower LeftRight `json:"lower"`
}
