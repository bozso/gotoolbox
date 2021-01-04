package path

import (

)

type GlobPattern struct {
    pattern string
    globbed []Valid
    err error
}

type GlobResult struct {}

func (g *GlobPattern) UnmarshalJSON(b []byte) (err error) {
    g.pattern = trim(b)
    g.globbed, g.err = New(g.pattern).Glob()
    return g.err
}

func (g GlobPattern) Glob() (v []Valid, err error) {
    return g.globbed, g.err
}

func (g GlobPattern) Into(in []From) (err error) {
    glob, err := g.Glob()
    if err != nil {
        return
    }

    for p := range glob {
        in

    }
}


type Globber struct {
    Valid    Valid           `json:"path"`
    Pattern  string          `json:"pattern"`
    Filterer FiltererPayload `json:"select"`
}

func (g Globber) Glob() (v []Valid, err error) {
    glob, err := g.Valid.Join(g.Pattern).Glob()
    if err != nil {
        return
    }
    
    v = make([]Valid, 0, 10)
    
    var keep bool
    for _, file := range glob {
        keep, err = g.Filterer.Filter(file)
        if err != nil {
            return
        }
        
        if keep {
            v = append(v, file)
        }
    }
    return
}

