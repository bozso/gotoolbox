package geometry

import (
    "github.com/bozso/gotoolbox/hash"
)

func (p Point2D) Hash(h hash.Hash) {
    hash.HashFloat64(p.X, h)
    hash.HashFloat64(p.Y, h)
}

func (ll LatLon) Hash(h hash.Hash) {
    ll.ToPoint().Hash(h)
}

func (m MinMaxFloat) Hash(h hash.Hash) {
    hash.HashFloat64(m.Min, h)
    hash.HashFloat64(m.Max, h)
}

func (r Region) Hash(h hash.Hash) {
    r.Min.Hash(h)
    r.Max.Hash(h)
}

func (r Rectangle) Hash(h hash.Hash) {
    r.X.Hash(h)
    r.Y.Hash(h)
}
