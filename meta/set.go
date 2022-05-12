package meta

type empty struct{}

var emptyVal = empty{}

type stringSet map[string]empty

type StringSet interface {
	HasString(s string) bool
	GetStrings() (strs []string)
}

type MutStringSet interface {
	StringSet
	SetString(s string)
	DelString(s string)
}

func (s StringSetImpl) HasString(str string) (ok bool) {
	_, ok = s.data[str]
	return
}

func (s StringSetImpl) GetStrings() (strs []string) {
	strs = make([]string, 0)
	for s := range s.data {
		strs = append(strs, s)
	}
	return
}

type StringSetImpl struct {
	data stringSet
}

func (s *StringSetImpl) SetString(str string) {
	s.data[str] = emptyVal
}

func (s *StringSetImpl) DelString(str string) {
	delete(s.data, str)
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
