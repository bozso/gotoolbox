package enum

import (
    "fmt"
    "strings"
)

type StringSet map[string]struct{}

func NewStringSet(s ...string) (ss StringSet) {
    ss = make(StringSet, len(s))
    for _, elem := range s {
        ss[elem] = struct{}{}
    }
    return
}

func (ss StringSet) Join(separator string) (s string) {
    var buf strings.Builder
    
    for key, _ := range ss {
        fmt.Fprintf(&buf, "%s%s", key, separator)
    }
    
    return buf.String()
}

func (ss StringSet) EnumType(name string) (t Type) {
    t.name, t.available = name, ss
    return
}

func (ss StringSet) Contains(s string) (b bool) {
    _, b = ss[s]
    return
}
