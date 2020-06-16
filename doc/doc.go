package doc

import (
    "time"
    "reflect"
    "github.com/bozso/gotoolbox/path"
)

type Doc struct {
    Title string
    date time.Time
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
    current int 
}

func (in Indices) Iter() (it Iter) {
    it.Indices, it.current = in, it.start
    return
}

func (it *Iter) Range() (reflect.Value, reflect.Value, bool) {
    for it.current < it.stop {
        curr := reflect.ValueOf(it.current)
        it.current += it.step
        
        return curr, curr, false 
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
    ind.stop, ind.start, ind.step = stop, 0, 1
    return
}
