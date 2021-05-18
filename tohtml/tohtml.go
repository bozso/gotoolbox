package tohtml

import (
	"sync"

	"github.com/bozso/gotoolbox/hash"

	tth "github.com/buildkite/terminal-to-html/v3"
)

type Renderer interface {
	Render([]byte) []byte
}

type cache map[hash.ID64][]byte

type Cache struct {
	cache
	hasher hash.Hasher
}

func NewCache(h hash.Hasher) (c Cache) {
	return Cache{
		hasher: h,
		cache:  make(cache),
	}
}

func (c *Cache) Render(b []byte) (out []byte) {
	c.hasher.Reset()
	c.hasher.Write(b)
	id := c.hasher.Sum64()

	out, ok := c.cache[id]
	if ok {
		return
	}

	out = tth.Render(b)
	c.cache[id] = out
	return
}

func (c Cache) WithMutex() (m MutexCache) {
	return MutexCache{
		Cache: c,
	}
}

type MutexCache struct {
	Cache
	mutex sync.Mutex
}

func (m *MutexCache) Render(b []byte) (out []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.Cache.Render(b)
}
