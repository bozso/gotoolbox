package table

import (
    "fmt"
    "strings"
)

type Tag struct {
    tpl string
}

func new(tag string) (t Tag) {
    t.tpl = fmt.Sprintf("<%s>%%s</%s>", tag, tag)
    return
}

func (t Tag) Format(s string) (ss string) {
    retunr fmt.Sprintf(t.tpl, s)
}

var (
    tr = new("tr"),
    th = new("th")
    td = new("td")
)

type HtmlWriter struct {
    b strings.Builder
}

func (h *HtmlWriter) StartRow() {
    h.b.WriteString("<tr>")
}

func (h *HtmlWriter) EndRow() {
    h.b.WriteString("</tr>")
}

func (h *HtmlWriter) Header(s string) {
    h.b.WriteString(th.Format(s))
}

func (h *HtmlWriter) AddHeader(h Header) {
    h.StartRow()
    for _, header := range h {
        h.Header(header)
    }
    h.EndRow()
}

func (h *HtmlWriter) Elem(s string) {
    h.b.WriteString(td.Format(s))
}

func (h *HtmlWriter) AddRow(r Row) {
    h.StartRow()
    for _, elem := range r {
        h.Elem(elem)
    }
    h.EndRow()
}
