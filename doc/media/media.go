package media

import (
    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/doc/result"
)

type Result struct {
    result.Status
    Media
}

type Media struct {
    path.ValidFile
}

func New(vf path.ValidFile) (r Result) {
    r.Media = Media{vf}
    return
}
