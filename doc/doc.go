package doc

import (
	"github.com/bozso/gotoolbox/path"
	"reflect"
	"time"
)

type Doc struct {
	Title string
	date  time.Time
}

func New() (d Doc) {
	return Doc{
		date: time.Now(),
	}
}

func (d Doc) WithTitle(title string) (D Doc) {
	d.Title = title
	return d
}

func (d Doc) Date() (s string) {
	const defaultFormat = "2006.01.01. 01:01"
	return d.DateWithFmt(defaultFormat)
}

func (d Doc) DateWithFmt(layout string) (s string) {
	return d.date.Format(layout)
}

func (_ Doc) NewPath(args ...string) (p path.Path) {
	return path.Joined(args...)
}

type Indices struct {
	start, stop, step int
}

func (in Indices) Start(start int) (ind Indices) {
	in.start = start
	return in
}

type Iter struct {
	Indices
	current, index int
}

func (in Indices) Iter() (it *Iter) {
	return &Iter{
		Indices: in,
		current: in.start,
		index:   0,
	}
}

func (it *Iter) Range() (reflect.Value, reflect.Value, bool) {
	for it.current < it.stop {
		it.current += it.step
		it.index += 1

		return reflect.ValueOf(it.index), reflect.ValueOf(it.current), false
	}

	return reflect.Value{}, reflect.Value{}, true
}

func (in Indices) Step(step int) (ind Indices) {
	in.step = step
	return in
}

func (in Indices) Stop(stop int) (ind Indices) {
	in.stop = stop
	return in
}

func (_ Doc) Indices(stop int) (ind Indices) {
	return Indices{
		stop:  stop,
		start: 0,
		step:  1,
	}
}
