package parsing

import ()

type FieldInfoMap map[string]Value

type StructInfo struct {
	FieldMap FieldInfoMap
}

func ParseStruct(g Getter, s *StructInfo) (err error) {
	for field, val := range s.FieldMap {
		if err = GetVal(g, field, val); err != nil {
			break
		}
	}

	return
}
