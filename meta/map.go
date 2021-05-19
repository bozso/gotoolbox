package meta

import (
	"reflect"
)

type MapKeysGet struct {
	Map  interface{}
	Keys []string
}

var MapKeys MapKeysGet

func (m *MapKeysGet) Call() (err error) {
	val := reflect.ValueOf(m.Map)
	if err = CheckKindOf(reflect.Map, val); err != nil {
		return
	}

	k := val.MapKeys()
	m.Keys = make([]string, len(k))

	for ii, key := range val.MapKeys() {
		if err = CheckKindOf(reflect.String, key); err != nil {
			return
		}
		/// TODO: check is key.String() would work
		m.Keys[ii] = key.Interface().(string)
	}

	return

}

func GetMapKeys(m interface{}) (keys []string, err error) {
	err = (&MapKeysGet{Map: m, Keys: keys}).Call()
	return
}
