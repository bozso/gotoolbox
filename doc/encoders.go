package doc

import (
    "bytes"
    "io"

    "github.com/bozso/gotoolbox/path"
)


type Encoder interface {
    Encode(dst, src []byte)
    EncodedLen(n int) int
}

type FileEncoder interface {
    EncodeFile(vf path.ValidFile) (s string, err error)
}

func EncodeFile(e Encoder, vf path.ValidFile) (b []byte, err error) {
    r, err := vf.Open()
    if err != nil {
        return
    }
    defer r.Close()

    b, err = EncodeTo(e, r)
    return
}

func EncodeTo(e Encoder, r io.Reader) (b []byte, err error) {
    var buf bytes.Buffer
    n, err := io.Copy(&buf, r)

    if err != nil {
        return
    }

    b = make([]byte, e.EncodedLen(int(n)))
    e.Encode(b, buf.Bytes())
    return
}

var noOpEncoder NoEncode

func NoOpEncoder() (f FileEncoder) {
    return &noOpEncoder
}

type NoEncode struct{}

func (_ NoEncode) EncodeFile(vf path.ValidFile) (s string, err error) {
    s = vf.String()
    return
}
