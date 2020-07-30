package server

import (
    "fmt"
    
    "github.com/tidwall/gjson"
)

type Field struct {
    name string
    gjson.Result
}

type Result struct {
    gjson.Result
}

func (r *Result) Set(s string) (err error) {
    r.Result = gjson.Parse(s)
    return nil
}

func (r Result) Get(field string) (f Field) {
    f.name, f.Result = field, r.Result.Get(field)
    return
}

func (f Field) EnsureType(t gjson.Type) (err error) {
    if t != f.Type {
        err = TypeError{expected: t, got: f.Type, field: f.name}
    }
    return
}

type TypeError struct {
    field string
    expected, got gjson.Type
}

func (e TypeError) Error() (s string) {
    return fmt.Sprintf("Expected type '%s' for field '%s', got '%s'",
        e.expected, e.field, e.got)
}

func (f Field) Exists() (err error) {
    if !f.Result.Exists() {
        err = Invalid{f.name}
    }
    return
}

type Invalid struct {
    field string
}

func (i Invalid) Error() (s string) {
    return fmt.Sprintf(
        "Expected field '%s' to be a valid or existing json value", i.field)
}

type FromJson interface {
    FromJson(gjson.Result) error
}

type ToJson interface {
    ToJson() (gjson.Result, error)
}
