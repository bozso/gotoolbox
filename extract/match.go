package extract

import (
	"io/fs"
	"path.filepath"
)

type Matcher struct {
	matches map[string]Creator
}

func NewMatcher(matches map[string]Creator) (m Matcher) {
	return Matcher{
		matches: matches,
	}
}

func (m Matcher) NewFs(path string) (fsys fs.FS, err error) {
	for key, val := range m.matches {
		matched, err := filepath.Match(path, key)

	}

}
