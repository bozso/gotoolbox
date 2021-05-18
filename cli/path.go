package cli

import (
	"fmt"
	"strings"

	"github.com/bozso/gotoolbox/path"
)

type Slice struct {
	content []string
}

func (s Slice) Iter() (sl []string) {
	return s.content
}

func (sl Slice) String() (out string) {
	if s := sl.content; s == nil {
		out = ""
	} else {
		out = fmt.Sprintf("%s\n", s)
	}

	return
}

func (sl *Slice) Set(s string) (err error) {
	err = nil
	slice := strings.Split(s, ",")

	if len(slice) == 0 {
		return
	}

	sl.content = slice
	return
}

func (sl Slice) Len() int {
	return len(sl.content)
}

type Paths []path.Path

func (sl Slice) ToPaths() (p Paths, err error) {
	p = make(Paths, sl.Len())
	s := sl.content

	for ii, _ := range s {
		p[ii] = path.New(s[ii])
	}
	return
}

type ZeroOrMore struct {
	Slice
}

func (z *ZeroOrMore) Set(s string) (err error) {
	return z.Slice.Set(s)
}

type OneOrMore struct {
	Slice
}

func (o *OneOrMore) Set(s string) (err error) {
	err = o.Slice.Set(s)
	if err != nil {
		return
	}

	if o.Len() == 0 {
		err = fmt.Errorf("expected at least one or more paths")
	}

	return
}
