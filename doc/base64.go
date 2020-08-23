package doc

import (
    "github.com/bozso/gotoolbox/path"
    "encoding/base64"
)

//const tpl = "data:{klass}/{mode};charset=utf-8;base64,{data}"
const tpl = "data:%s/%s;charset=utf-8;base64,%s"

func Encode(p path.ValidFile) (s string, err error) {
    content, err := p.ReadAll()
    if err != nil {
        return
    }
    
    s = base64.StdEncoding.EncodeToString(content)
    return
}
