package path

import (
)

type IndexedFrom interface {
    GetFrom(ii int) From
}

type From interface {
    FromPath(p Pather) error
}

