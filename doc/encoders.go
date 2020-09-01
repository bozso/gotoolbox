package doc

import (
    "fmt"
    "bytes"
    "io"
    "encoding/base64"
    
    "github.com/bozso/gotoolbox/path"
)


type Encoder interface {
    Encode(dst, src []byte)
}

type (
    StringSet map[string]struct{}
    typeToExtension map[string]StringSet
)

func NewStringSet(s ...string) (ss StringSet) {
    ss = make(StringSet, len(s))
    for _, elem := range s {
        ss[elem] = struct{}{}
    }
    return
}

func (ss StringSet) Contains(s string) (b bool) {
    _, b = ss[s]
    return
}

var class2ext = typeToExtension{
    "text" : NewStringSet("js", "css"),
    "video" : NewStringSet("mp4"),
    "image": NewStringSet("png", "jpg"),
}

func ExtensionToType(ext string) (extType string, err error) {
    for key, val := range class2ext {
        if val.Contains(ext) {
            return key, nil
        }
    }
    
    err = fmt.Errorf("could not find a matching class for extension: '%s'",
        ext)
    return
}

var convert = map[string]string{
    "js": "javascript",
}

type HtmlEncoder struct {
    encoder Encoder
}

func (h *HtmlEncoder) set() {
    if h.encoder == nil {
        h.encoder = base64.StdEncoding
    }
}

func (h HtmlEncoder) EncodeFile(vf path.ValidFile) (b []byte, err error) {
    h.set()

    ext := vf.Ext()
    
    extType, err := ExtensionToType(ext)
    if err != nil {
        return
    }
    
    buf := &bytes.Buffer{}
    fmt.Fprintf(buf, "data:%s/%s;charset=utf-8;base64,", extType, ext)
    
    out, err := EncodeFile(h.encoder, vf)
    if err != nil {
        return
    }
    
    buf.Write(out)
    return buf.Bytes(), nil
}

func EncodeFile(e Encoder, vf path.ValidFile) (b []byte, err error) {
    r, err := vf.Open()
    if err != nil {
        return
    }
    
    err = EncodeTo(e, r, b)
    return
}

func EncodeTo(e Encoder, r io.Reader, b []byte) (err error) {
    buf := &bytes.Buffer{}
    n, err := io.Copy(buf, r)
    
    if err != nil {
        return
    }
    
    
    e.Encode(b, buf.Bytes())
    return    
}


