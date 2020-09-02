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

func (ss StringSet) Contains(s string) (b bool) {
    _, b = ss[s]
    return
}

type Type string

func (t Type) UnknownElement(set StringSet, value string) (u UnknownElement) {
    u.EnumType, u.Value, u.Set = string(t), value, set
    return
}

type UnknownElement struct {
    Set StringSet
    Value, EnumType string
}

func (u UnknownElement) Error() (s string) {
    return fmt.Sprintf("unknown value '%s' for %s, choose from %s", u.Value,
        u.EnumType, u.Set.Join(","))
}
