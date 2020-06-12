package doc

import (

)

type Getter interface {
    Get() error
}

type Status struct {
    Err error
}

type CaptureFn func()

func (s Status) Use(err *error) CaptureFn {
    return func() {
        s.Err = *err
    }
}

type Capture struct {
    s *Status
    err *error
}

func (c Capture) Set() {
    c.s.Err = *c.err
}

func (s Status) IsErr() (b bool) {
    return s.Err != nil
}

func (s *Status) Set(err error) {
    s.Err = err
}

func (s Status) Get() (err error) {
    return s.Err
}

func (s *Status) From(g Getter) (b bool) {
    if s.Err != nil {
        return true
    }
    
    err := g.Get()

    if err != nil {
        s.Err = err
        b = true
    } else {
        b = false
    }
    
    return
}
