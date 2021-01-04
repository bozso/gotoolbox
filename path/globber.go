package path

import (

)

type GlobPattern struct {
    pattern string
    result GlobResult
    err error
}

func (g *GlobPattern) UnmarshalJSON(b []byte) (err error) {
    g.pattern = trim(b)
    globbed, err := New(g.pattern).Glob()
    g.result, g.err = GlobResult{globbed}, err
    return err
}

func (g GlobPattern) Glob() (gr GlobResult, err error) {
    return g.result, g.err
}

type GlobResult struct {
    glob []Valid
}

func (g GlobResult) Len() (n int) {
    return len(g.glob)
}

func (g GlobResult) Into(in []From) (err error) {
    for ii, p := range g.glob {
        err = in[ii].FromPath(p)
        if err != nil {
            break
        }
    }
    return
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

