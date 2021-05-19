package path

import (
	"github.com/bozso/gotoolbox/errors"
)

type GlobPatterns map[string]GlobPattern

func (g GlobPatterns) Get(key string, ext Extension) (f []ValidFile, err error) {
	p, ok := g[key]
	if !ok {
		err = errors.KeyNotFound(key)
		return
	}

	result, err := p.Glob()
	if err != nil {
		return
	}

	files := ext.Files(result.Len())
	err = result.Into(files)
	f = files.Files()
	return
}
