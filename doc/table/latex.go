package table

import (
    "strings"
)

// TODO: implement this

type LatexWriter struct {
    strings.Builder
}

func (h *LatexWriter) StartRow() {
}

func (h *LatexWriter) EndRow() {
}

func (h *LatexWriter) Header(s string) {
}

func (h *LatexWriter) AddHeader(head Header) {
}

func (h *LatexWriter) Elem(s string) {
}

func (h *LatexWriter) AddRow(r Row) {
}
