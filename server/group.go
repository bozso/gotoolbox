package server

import (
    "github.com/bozso/gotoolbox/errors"
)

const groupName errors.Name = "group"

type groups map[string]Methods

type Groups struct {
    groups
}

func NewGroupsWithCap(cap int) (g Groups) {
    m.groups = make(groups, 0, cap)
    m.available = groupName.NewAvailable()
    return
}

func NewGroups() (g Groups) {
    return NewGroupsWithCap(defaultCap)
}

func (g *Groups) Add(name string, m Methods) {
    m.groups[name] = m
    m.available.Add(name)
}

func (g Groups) Handle(group, method string, in, out []byte) (err error) {
    g, ok := m.groups[group]
    if !ok {
        return g.available.NotFound(group)
    }
    
    err = g.Handle(merhod, in, out)
    return
}
