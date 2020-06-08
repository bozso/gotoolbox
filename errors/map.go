package errors

import (
    "fmt"
)

type KeyError struct {
    key string
}

func KeyNotFound(key string) (k KeyError) {
    return KeyError{key}
}

func (e KeyError) Error() (s string) {
    return fmt.Sprintf("key '%s' not found in map")
}
