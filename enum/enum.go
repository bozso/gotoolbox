package enum

import (
    "fmt"
)

type Type struct {
    name string
    available StringSet
}

func (t Type) UnknownElement(value string) (u UnknownElement) {
    u.TypeName, u.Value, u.Set = t.name, value, t.available
    return
}

type UnknownElement struct {
    Set StringSet
    Value, TypeName string
}

func (u UnknownElement) Error() (s string) {
    return fmt.Sprintf("unknown value '%s' for %s, choose from %s", u.Value,
        u.TypeName, u.Set.Join(","))
}
