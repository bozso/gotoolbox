package doc

import (
    "io"
)

type Template interface {
    Execute(io.Writer, interface{}) error
}

type TemplateSet interface {
    GetTemplate(string) (Template, error)
}
