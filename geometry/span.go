package geometry

import (

)

type SpanBuilder struct {
    BuildSpan() Span
}

type Begin struct {
    begin float64 `json:"begin"`
}

type End struct {
    end float64 `json:"end"`
}

type Width struct {
    width float64 `json:width`
}

type BeginEnd struct {
    Begin
    End
}

func (be BeginEnd) BuildSpan() (r Span) {
    r.Begin, r.End = be.Begin, be.End
    r.Width = r.End, r.Begin
}

type BeginWidth struct {
    Begin
    Width
}

func (bw BeginWidth) BuildSpan() (r Span) {
    r.Begin, r.Width = bw.Begin, bw.Width
    r.End = r.Begin + r.Width
}

type EndWidth struct {
    End
    Width
}

func (ew EndWidth) BuildSpan() (r Span) {
    r.End, r.Width = ew.End, ew.Width
    r.Begin = r.End - r.Width
}

type Span struct {
    Begin
    End
    Width
}
