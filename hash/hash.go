package hash

import (
    "io"
    "crypto"
    stdHash "hash"
)

type Hashable interface {
    Hash(h *stdHash.Hash)
}

type Hasher struct {
    stdHash.Hash
}

func NewHasher(mode crypto.Hash) (h Hasher) {
    h.Hash = mode.New()
    return
}

func (h *Hasher) Hash(hsb Hashable) (b []byte) {
    h.Reset()
    h.Append(hsb)
    
    return h.Sum(nil)
}

func (h *Hasher) Append(hsb Hashable) {
    hsb.Hash(h)
}
