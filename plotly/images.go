package plotly

import (
    "github.com/bozso/emath/geometry"
    "github.com/bozso/gotoolbox/path"
)

type Image struct {
    File path.ValidFile    `json:"file"`
    Extent geometry.Region `json:"extent"`
    Title string           `json:"title"`
}

type Snapshots struct {
    Colorbar []path.ValidFile
    Images []Image
}
