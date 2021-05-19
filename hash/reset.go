package hash

import (
	"hash"
)

type Resetter struct {
	h hash.Hash64
}

func NewResetter(h hash.Hash64) (r Resetter) {
	return Resetter{
		h: h,
	}
}

func (r *Resetter) UseHash(fn UseFn) (err error) {
	r.h.Reset()
	return fn(r.h)
}
