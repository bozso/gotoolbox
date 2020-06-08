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


const splitErr errors.String = "could not split string '%s' into fields"

func NewFields(s string) (sp Parser, err error) {
    sp, err = FromSlice(strings.Fields(s))
    if err != nil {
        splitErr.WrapFmt(err, s)
    }
    return
}

const parseErr errors.String = "could not parse string '%s' with separator '%s'"

func New(s, sep string) (sp Parser, err error) {
    sp, err = FromSlice(strings.Split(s, sep))
    if err != nil {
        parseErr.WrapFmt(err, s, sep)
    }
    
    return
}

func (sp Parser) Len() (e errors.Bound) {
    return errors.Bound(sp.len)
}

func (sp Parser) Idx(idx int) (s string, err error) {
    if err = sp.Len().IsOutOfBounds(idx); err != nil {
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
