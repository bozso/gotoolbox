package doc

import (
    "fmt"
    "strings"
    "encoding/base64"

    "github.com/bozso/gotoolbox/path"
)

type HtmlEncoder struct {
    encoder Encoder
}

func NewHtmlEncoder(enc Encoder) (h HtmlEncoder) {
    h.encoder = enc
    return
}

func (h HtmlEncoder) EncodeFile(vf path.ValidFile) (s string, err error) {
    ext := vf.Extension()[1:]

    extType, err := ExtensionToType(string(ext))
    if err != nil {
        return
    }

    var buf strings.Builder
    fmt.Fprintf(&buf, "data:%s/%s;charset=utf-8;base64,", extType, ext)

    out, err := EncodeFile(h.encoder, vf)
    if err != nil {
        return
    }

    buf.Write(out)
    return buf.String(), nil
}

type HtmlEncoderPlotly struct {
    encoder Encoder
}

func NewHtmlPlotlyEncoder(enc Encoder) (p HtmlEncoderPlotly) {
    p.encoder = enc
    return
}

func (p HtmlEncoderPlotly) EncodeFile(vf path.ValidFile) (s string, err error) {
    ext := vf.Extension()[1:]

    var buf strings.Builder
    fmt.Fprintf(&buf, "data:image/%s;base64,", ext)

    out, err := EncodeFile(p.encoder, vf)
    if err != nil {
        return
    }

    buf.Write(out)
    return buf.String(), nil
}

var Base64 = NewHtmlEncoder(base64.StdEncoding)
var PlotlyBase64 = NewHtmlPlotlyEncoder(base64.StdEncoding)
