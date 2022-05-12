package meta

type empty struct{}

var emptyVal = empty{}

type stringSet map[string]empty

type StringSet interface {
	HasString(s string) bool
}

type MutStringSet interface {
	StringSet
	SetString(s string)
}

type StringSetImpl struct {
	data stringSet
}

func (s *StringSetImpl) SetString(str string) {
	s.data[str] = emptyVal
}

func (s StringSetImpl) HasString(str string) (ok bool) {
	_, ok = s.data[str]
	return
}

func EmptyStringSet() (ss StringSetImpl) {
	return StringSetImpl{
		data: make(stringSet),
	}
}

func AddSlice(ms MutStringSet, strs []string) {
	for _, s := range strs {
		ms.SetString(s)
	}
}

func NewStringSet(strs []string) (ss StringSetImpl) {
	ss = EmptyStringSet()
	AddSlice(&ss, strs)
	return ss
}
