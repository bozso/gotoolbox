package doc

import (
    "bytes"
    "io"
    
    "github.com/bozso/gotoolbox/path"
)


type Encoder interface {
    Encode(dst, src []byte)
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
    
    
    err = EncodeTo(e, r, b)
    return
}

func EncodeTo(e Encoder, r io.Reader, b []byte) (err error) {
    var buf bytes.Buffer
    _, err = io.Copy(&buf, r)
    
    if err != nil {
        return
    }
    
    
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
