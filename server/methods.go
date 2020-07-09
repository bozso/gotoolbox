package server

import (
    "fmt"
    
    "github.com/bozso/gotoolbox/errors"
)

type methods map[string]Handler 

type Methods struct {
    methods
    available errors.Available
}

const defaultCap = 4

const methodsName errors.Name = "method"

func NewMethodsWithCap(cap int) (m Methods) {
    m.methods = make(methods, 0, cap)
    m.available = methodsName.NewAvailable()
    return
}

func NewMethods() (m Methods) {
    return NewMethodsWithCap(defaultCap)
}

func (m *Methods) Add(name string, h Handler) {
    m.methods[name] = h
    m.available.Add(name)
}

func (m Methods) Handle(name string, in, out []byte) (err error) {
    h, ok := m.methods[name]
    if !ok {
        return m.available.NotFound(name)
    }
    
    err = h.Handle(in, out)
    return
}
