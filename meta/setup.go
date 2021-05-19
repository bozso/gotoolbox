package meta

import (
	"sync"
)

type Callable interface {
	Call() error
}

type Startup struct {
	once sync.Once
}

type CallFunc func() error

func (c CallFunc) Call() (err error) {
	return c()
}

type Empty struct{}

func (s *Startup) Do(call Callable) (e Empty) {
	s.once.Do(func() {
		if err := call.Call(); err != nil {
			panic(err)
		}
	})
	return
}
