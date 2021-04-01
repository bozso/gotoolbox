package meta

import (
    "fmt"
    "reflect"
)

func CheckKindOf(kind reflect.Kind, val reflect.Value) (err error) {
    return CheckKind(kind, val.Kind())
}

func CheckKind(expected, got reflect.Kind) (err error) {
    err = nil
    if expected != got {
        return KindMismatch {
            Expected: expected,
            Got: got,
        }
    }
    return
}

type KindMismatch struct {
    Expected, Got reflect.Kind
}

func (k KindMismatch) Error() (s string) {
    return fmt.Sprintf("expected kind '%s', got '%s'", 
        k.Expected.String(), k.Got.String())
}
