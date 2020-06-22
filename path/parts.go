package path

import (
    "encoding/json"
)

type Parts struct {
    parts []string
}

func (d *Parts) UnmarshalJSON(b []byte) (err error) {
    err = json.Unmarshal(b, &d.parts)
    return
}

func (p Parts) ToPath() (pp Path) {
    pp = Joined(p.parts...)
    return
}
