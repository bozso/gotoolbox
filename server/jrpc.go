package server

import (
    "fmt"
    "encoding/json"
)

const version Version = "2.0"


type Base struct {
    jsonrpc string `json:"jsonrpc"`
    id      string `json:"id,omitempty"`
}

type Request struct {
    Base
    method string `json:"method"`
    params []byte `json:"params,omitempty"`
}

type Success struct {
    Base
    result []byte `json:"result"`
}

func (s Success) Respond(w io.Writer) (err error) {
    b, err := json.Marshal(f)
    if err != nil {
        return
    }
    
    _, err = w.Write(b)
    return
}

type Failure struct {
    Base
    err string `json:"error"`
}

func (f Failure) Respond(w io.Writer) (err error) {
    b, err := json.Marshal(f)
    if err != nil {
        return
    }
    
    _, err = w.Write(b)
    return
}

type Responder interface {
    Respond(w io.Writer) (err error)
}

type JSONRpc struct {
    Groups
    //paramName string
}

var jsonType = []byte{"application/json"}

func (j JSONRpc) HandleImpl(ctx *fasthttp.RequestCtx) (err error) {
    cReq := ctx.Request
    
    if cReq.Header.ContentType != jsonType {
        err = WrongContentType{}
        return
    }
    
    var req Request
    if err := json.Unmarshal(cReq.Body(), &req); err != nil {
        return
    }
    
    var out []byte
    
    err = j.Groups.Handle(group, method, req.params, out)
    
    if err != nil {
        
    }
    
    return
}

type Handler interface {
    Handle(in, out []byte) (err error)
}
