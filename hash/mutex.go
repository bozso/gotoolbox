package hash

import (
	"sync"
)

type Mutexed struct {
	mutex   sync.Mutex
	service Service
}

func NewMutexed(s Service) (m Mutexed) {
	return Mutexed{
		service: s,
	}
}

func (m *Mutexed) UseHash(fn UseFn) (err error) {
	m.mutex.Lock()
	m.service.UseHash(fn)
	m.mutex.Unlock()
	return
}
