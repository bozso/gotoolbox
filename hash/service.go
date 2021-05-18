package hash

import (
	"hash"
)

type Service interface {
	UseHasher(func(hash.Hash64) error) error
}
