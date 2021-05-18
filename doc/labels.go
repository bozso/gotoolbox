package doc

import (
	"fmt"
)

type String string
type formatted string

func (s String) Format(args ...interface{}) (ss String) {
	return String(fmt.Sprintf(string(s), args...))
}

func (s String) Error() (ss string) {
	return string(s)
}

const KeyError String = "key '%s' not found in LabelIndex"

type dict map[string]int

type LabelIndex struct {
	dict
	idx int
}

func NewLabels() (li LabelIndex) {
	return LabelIndex{
		dict: make(map[string]int),
		idx:  0,
	}
}

func (li LabelIndex) Iter() (d dict) {
	return li.dict
}

func (li *LabelIndex) AddIdx(key string, idx int) {
	li.dict[key] = idx
}

func (li *LabelIndex) Add(key string) {
	li.dict[key] = li.idx
	li.idx += 1
}

func (li LabelIndex) Get(key string) (idx int, err error) {
	idx, ok := li.dict[key]
	if !ok {
		err = KeyError.Format(key)
	}
	return
}
