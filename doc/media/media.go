package media

import (
    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/doc"
)

type Result struct {
    doc.Status
    Media
}

type Media struct {
    path.ValidFile
}

func New(vf path.ValidFile) (r Result) {
    r.Media = Media{vf}
    return
}
