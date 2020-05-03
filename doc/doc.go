package doc

import (
    "time"
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

func (d Doc) NewPath(args ...string) (p path.Path) {
    return path.Joined(args...)
}
