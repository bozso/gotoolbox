package hash

import (
	"hash"
	"sync"
)

type PoolService struct {
	pool sync.Pool
}

func NewPool(c Creator) (p PoolService) {
	return PoolService{
		pool: sync.Pool{
			New: func() interface{} {
				return c.CreateHash()
			},
		},
	}
}

func (p PoolService) Get() (h hash.Hash64) {
	return p.pool.Get().(hash.Hash64)
}

func (p PoolService) Put(h hash.Hash64) {
	p.pool.Put(h)
}

func (p PoolService) UseHash(fn UseFn) (err error) {
	h := p.Get()
	err = fn(h)
	p.Put(h)

	return
}
