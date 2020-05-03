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
    return fmt.Sprintf(t.tpl, s)
}

var tr, th, td = new("tr"), new("th"), new("td")

type HtmlWriter struct {
    strings.Builder
}

func (h *HtmlWriter) StartRow() {
    h.Builder.WriteString("<tr>")
}

func (h *HtmlWriter) EndRow() {
    h.Builder.WriteString("</tr>")
}

func (h *HtmlWriter) Header(s string) {
    h.Builder.WriteString(th.Format(s))
}

func (h *HtmlWriter) AddHeader(head Header) {
    h.StartRow()
    for _, header := range head {
        h.Header(header)
    }
    h.EndRow()
}

func (h *HtmlWriter) Elem(s string) {
    h.Builder.WriteString(td.Format(s))
}

func (h *HtmlWriter) AddRow(r Row) {
    h.StartRow()
    for _, elem := range r {
        h.Elem(elem)
    }
    h.EndRow()
}
