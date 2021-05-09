package path

import (
	"fmt"
	"os"
	"time"

	"github.com/bozso/gotoolbox/errors"
)

type Valid struct {
	Path
}

func (v *Valid) FromPath(p Pather) (err error) {
	*v, err = p.AsPath().ToValid()
	return
}

func (v *Valid) Set(s string) (err error) {
	const name errors.NotEmpty = "valid path"

	if err = name.Check(s); err != nil {
		return
	}

	*v, err = New(s).ToValid()
	return
}

func (v *Valid) UnmarshalJSON(b []byte) (err error) {
	err = v.Set(trim(b))
	return
}

func (v Valid) ToFile() (vf ValidFile, err error) {
	isDir, err := v.IsDir()
	if err != nil {
		return
	}

	if isDir {
		err = fmt.Errorf("path '%s' is not a file, but a directory", v)
	}

	vf.Path = v.Path
	return
}

func (vp Valid) Stat() (fi os.FileInfo, err error) {
	fi, err = os.Lstat(vp.String())

	if err != nil {
		err = vp.Fail(OpStat, err)
	}

	return
}

func (vp Valid) ModTime() (t time.Time, err error) {
	fi, err := vp.Stat()
	if err != nil {
		return
	}

	t = fi.ModTime()
	return
}

func (v Valid) OlderThan(vp Valid) (b bool, err error) {
	t1, err := v.ModTime()
	if err != nil {
		return
	}

	t2, err := v.ModTime()
	if err != nil {
		return
	}

	b = t1.After(t2)
	return
}

func (v Valid) YoungerThan(vp Valid) (b bool, err error) {
	b, err = v.OlderThan(vp)
	b = !b
	return
}

func (vp Valid) IsDir() (b bool, err error) {
	b = false
	fi, err := vp.Stat()
	if err != nil {
		return
	}

	b = fi.IsDir()
	return
}

func (vp Valid) Open() (of *os.File, err error) {
	of, err = os.Open(vp.String())

	if err != nil {
		err = vp.Fail(OpOpen, err)
	}

	return
}

func (vp Valid) Remove() (err error) {
	return os.Remove(vp.String())
}

func (vp Valid) Rename(target fmt.Stringer) (dst Valid, err error) {
	s1, s2 := vp.String(), target.String()

	if err = os.Rename(s1, s2); err != nil {
		err = errors.WrapFmt(err, "failed to move '%s' to '%s'", s1, s2)
		return
	}

	dst, err = New(s2).ToValid()
	return
}

func (vp Valid) Move(dir Dir) (dst Valid, err error) {
	_dst, err := dir.Join(vp.Base().String()).Abs()
	if err != nil {
		return
	}

	dst, err = vp.Rename(_dst)
	return
}

type validPathOperation int

const (
	OpOpen validPathOperation = iota
)

func (op validPathOperation) Fmt(p Path) (s string) {
	ps := p.GetPath()

	switch op {
	case OpOpen:
		s = fmt.Sprintf("failed to open path '%s'", ps)
	}
	return
}
