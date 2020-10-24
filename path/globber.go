package path

import (

)

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

