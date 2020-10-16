package path

import (
    
)

type Glob func() ([]Valid, error)
type Filter func(Pather) (bool, error)
type EnumFilter func(int, Pather) (bool, error)
type Transform func(Pather) (Pather, error)

type Filters []Filter
type Transforms []Transform

func SelectAll(_ Pather) bool {
    return true
}

