package path

import (
    
)

type Filter func(Path) bool
type Transform func(Path) Path

type Filters []Filter
type Transforms []Transform
