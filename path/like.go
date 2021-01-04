package path

import (
    "fmt"
)

type Convertable interface {
    ToValid() (Valid, error)
    ToFile() (File, error)
    ToDir() (Dir, error)
    ToValidFile() (ValidFile, error)
}

type Operations interface {
    fmt.Stringer
    Abs() (Like, error)
    Join(...string) Like
    Extension() Extension
    Exists() (bool, error)
    Glob() ([]Valid, error)
}

type Like interface {
    Convertable
    Operations
}
