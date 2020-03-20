package splitted

import (
    "fmt"
    "strings"
    "strconv"

    "github.com/bozso/gotoolbox/errors"
)

type Parser struct {
    split []string
    len int
}

func FromSlice(slice []string) (sp Parser, err error) {
    sp.len = len(slice)

    if sp.len == 0 {
        err = fmt.Errorf("got empty slice")
        return
    }
    
    sp.split = slice
    return
}

func NewFields(s string) (sp Parser, err error) {
    sp, err = FromSlice(strings.Fields(s))
    if err != nil {
        errors.WrapFmt(err, "could not split string '%s' into fields", s)
    }
    return
}

func New(s, sep string) (sp Parser, err error) {
    sp, err = FromSlice(strings.Split(s, sep))
    if err != nil {
        errors.WrapFmt(err,
            "could not parse string '%s' with separator '%s'", s, sep)
    }
    
    return
}

func (sp Parser) Len() int {
    return sp.len
}

func (sp Parser) Idx(idx int) (s string, err error) {
    if err = errors.IsOutOfBounds(idx, sp.len); err != nil {
        return
    }
    
    return sp.split[idx], nil
}

func (sp Parser) Int(idx int) (ii int, err error) {

    s, err := sp.Idx(idx)
    
    if err != nil {
        return
    }
    
    ii, err = strconv.Atoi(s)
    return
}

func (sp Parser) Float(idx int) (ff float64, err error) {
    s, err := sp.Idx(idx)
    if err != nil { return }
    
    ff, err = strconv.ParseFloat(s, 64)
    
    return
}
