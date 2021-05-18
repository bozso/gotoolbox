package hash

import (
	"hash"
)

type UseFn func(hash.Hash64) error

type Service interface {
	UseHash(UseFn) error
}

func Bytes(s Service, hs Hashable) (b []byte) {
	s.UseHash(func(h hash.Hash64) (err error) {
		hs.Hash(h)
		b = h.Sum(nil)
		return nil
	})
	return b
}

func Sum64(s Service, hs Hashable) (id ID64) {
	s.UseHash(func(h hash.Hash64) (err error) {
		hs.Hash(h)
		id = ID64(h.Sum64())
		return nil
	})
	return id
}

type Creator interface {
	CreateHash() hash.Hash64
}
