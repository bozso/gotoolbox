package path

import (
    "fmt"
)

type Extension string

func (e Extension) Check(v ValidFile) (b bool) {
    return v.Ext() == string(e)
}

func (e Extension) MustHave(v ValidFile) (err error) {
    if ext := v.Ext(); ext != string(e) {
        err = WrongExtension{
            Expected: e,
            Got: ext,
            File: v,
        }
    }
    return
}

type WrongExtension struct {
    Expected Extension
    Got string
    File ValidFile
}

func (e WrongExtension) Error() (s string) {
    return fmt.Sprintf(
        "expected extension '%s' got '%s' for path '%s'", e.Expected,
        e.Got, e.File)
}
