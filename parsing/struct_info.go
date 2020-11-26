package parsing

import (

)

type FieldInfoMap map[string]Value

type StructInfo struct {
    FieldMap FieldInfoMap
}

func (g GetterParser) ParseStruct(s *StructInfo) (err error) {
    for field, val := range s.FieldMap {
        if err = g.GetVal(field, val); err != nil {
            break
        }
    }

    return
}
