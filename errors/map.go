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

func KeyNotFoundWith(key interface{}) (k KeyError) {
    str, ok := key.(fmt.Stringer)
    
    var s string
    
    if ok {
        s = str.String()
    } else {
        s = fmt.Sprintf("%#v", key)
    }
    
    return KeyError{s}
}
