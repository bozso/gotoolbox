package result

import (

)

type Getter interface {
    Get() error
}

type Status struct {
    Err error
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

func (s *Status) From(g Getter) {
    s.Err = g.Get()
}
