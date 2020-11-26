package parser

import (

)

type Parser struct {
    getter Getter
}

func New(g Getter) (p Parser) {
    return Parser {
        getter: g,
    }
}

func (p Parser) Get(key string) (s string, b bool) {
    return Get(p.getter, key)
}

func (p Parser) MustGet(key string) (s string, err error) {
    return MustGet(p.getter, key)
}

func (p Parser) Int(key string) (ii int, b bool) {
    return GetInt(p.getter, key)
}
