package main

import (
    "testing"
    "sync"
)

type Mutable struct {
    ii int
    mutex sync.RWMutex
}

func withLock(m *sync.RWMutex, fn func()) {
    m.Lock()
    fn()
    m.Unlock()
}

func withRLock(m *sync.RWMutex, fn func()) {
    m.RLock()
    fn()
    m.RUnlock()
}

func BenchmarkManualWrite(b *testing.B) {
    m := Mutable{ ii: 0 }

    for ii := 0; ii < b.N; ii++ {
        m.mutex.Lock()
        m.ii++
        m.mutex.Unlock()
    }
}

func BenchmarkClosureWrite(b *testing.B) {
    m := Mutable{ ii: 0 }

    for ii := 0; ii < b.N; ii++ {
        withLock(&m.mutex, func() {
            m.ii++
        })
    }
}

func BenchmarkManualRead(b *testing.B) {
    m := Mutable{ ii: 0 }
    idx := 0

    for ii := 0; ii < b.N; ii++ {
        m.mutex.RLock()
        idx = m.ii
        m.mutex.RUnlock()
    }
    _ = idx
}

func BenchmarkClosureRead(b *testing.B) {
    m := Mutable{ ii: 0 }

    idx := 0
    for ii := 0; ii < b.N; ii++ {
        withRLock(&m.mutex, func() {
            idx = m.ii
        })
    }
    _ = idx
}
