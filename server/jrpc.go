package server

import (
    "fmt"
    
    "github.com/valyala/fasthttp"
    "github.com/tidwall/gjson"
)

type Base struct {
    jsonrpc string
    id gjson.Result
}

type Version string

const version Version = "2.0"

type VersionMismatch struct {
    expected Version
    got string
}

func (v VersionMismatch) Error() (s string) {
    return fmt.Sprintf("Expected version string to be '%s', got '%s'",
        v.expected, v.got)
}

func (v Version) Validate(s string) (err error) {
    if s != string(v) {
        err = VersionMismatch{expected: v, got: s}
    }
    return
}

func (b *Base) FromJson(r Result) (err error) {
    j := r.Get("jsonrpc")
    
    if err = j.Exists(); err != nil {
        return
    }
    
    jsonrpc := j.String()
    
    if err = version.Validate(jsonrpc); err != nil {
        return err
    }
    
    j  = r.Get("id")

    if err = j.Exists(); err != nil {
        return
    }
    
    b.id = j.Result
    return
}

type Request struct {
    Base
    method string
    params Field
}

func (req *Request) FromJson(r Result) (err error) {
    if err = req.Base.FromJson(r); err != nil {
        return
    }
    
    j := r.Get("method")
    if err = j.Exists(); err != nil {
        return
    }
    req.method = j.String()
    
    j = r.Get("params")
    if err = j.Exists(); err != nil {
        return
    }
    
    req.params = j
    return
}

type JSONRpc struct {
    methods map[string]Handler
    paramName string
}

func (j JSONRpc) Handle(ctx *fasthttp.RequestCtx) {
    
}

type Handler interface {
    Handle([]byte) (Marshaller, error)
}

type Marshaller interface {
    Marshall() ([]byte, error)
}
