package doc

import (
    "sync"
    
    "github.com/bozso/gotoolbox/path"    
)

type EncodedFiles map[path.ValidFile]string

type EncodeCache struct {
    db EncodedFiles
    mutex sync.RWMutex
}

func NewEncodeCache() (e EncodeCache) {
    e.db = make(EncodedFiles)
    return
}

func (e *EncodeCache) Get(p path.ValidFile) (s string, ok bool) {
    e.mutex.RLock()
    s, ok = e.db[p]
    e.mutex.RUnlock()
    return
}

func (e *EncodeCache) Set(p path.ValidFile, s string) {
    e.mutex.Lock()
    e.db[p] = s
    e.mutex.Unlock()
}

func (e EncodeCache) WithEncoder(fe FileEncoder) (cf CachedFileEncoder) {
    cf.fileEncoder, cf.cache = fe, e
    return
}

type CachedFileEncoder struct {
    cache EncodeCache
    fileEncoder FileEncoder
}

func (c CachedFileEncoder) EncodeFile(vf path.ValidFile) (s string, err error) {
    s, ok := c.cache.Get(vf)
    if ok {
        return
    }
    
    s, err = c.fileEncoder.EncodeFile(vf)
    if err != nil {
        return
    }
    
    c.cache.Set(vf, s)
    return
}

