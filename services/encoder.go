package services

import (
    //"fmt"
    "sync"
    "strings"
    "net/http"
    "encoding/base32"
    "encoding/base64"

    "github.com/bozso/gotoolbox/enum"
    "github.com/bozso/gotoolbox/doc"
    "github.com/bozso/gotoolbox/hash"
)

type ID hash.ID64

type Encoder interface {
    New(r *http.Request, args *string, result *ID) (err error)
    Delete(r *http.Request, args *ID, result *Empty) (err error)
    EncodeFile(r *http.Request, args *Path, result *string) (err error)
}

type EncoderService struct {
    db map[ID]doc.Encoder
    mutex sync.RWMutex
}

type EncoderType int

const (
    Base32 EncoderType = iota
    Base64
)

var encoderType = enum.NewStringSet("Base32", "Base64").EnumType("EncoderType")

func (et EncoderType) String() (s string) {
    switch et {
    case Base32:
        s = "Base32"
    case Base64:
        s = "Base64"
    default:
        s = "Unknown"
    }
    return
}

func (et *EncoderType) Set(s string) (err error) {
    switch strings.ToLower(s) {
    case "base32":
        *et = Base32
    case "base64":
        *et = Base64
    default:
        err = encoderType.UnknownElement(s)
    }
    
    return
}

func (et EncoderType) New() (e doc.Encoder) {
    switch et {
    case Base32:
        e = base32.StdEncoding
    case Base64:        
        e = base64.StdEncoding
    }
    return
}

func (e *EncoderService) Get(id ID) (enc doc.Encoder, ok bool) {
    e.mutex.RLock()
    enc, ok = e.db[id]
    e.mutex.RUnlock()

    return
}

func (e *EncoderService) Set(id ID, enc doc.Encoder) {
    e.mutex.Lock()
    e.db[id] = enc
    e.mutex.Unlock()
}

func (e *EncoderService) new(et EncoderType) (id ID) {
    id = ID(et)
    _, ok := e.Get(id)
    
    if ok {
        return
    }
    
    e.Set(id, et.New())
    return
}

func (e *EncoderService) New(r *http.Request, args *string, result *ID) (err error) {
    var et EncoderType
    if err = et.Set(*args); err != nil {
        return
    }
    
    *result = e.new(et)
    
    
    return
}

func (e *EncoderService) Delete(r *http.Request, args *ID, result *Empty) (err error) {
    e.mutex.Lock()
    delete(e.db, *args)
    e.mutex.Unlock()

    return
}

// TODO: implement
func (e *EncoderService) EncodeFile(r *http.Request, args *Path, result *string) (err error) {
    return
}
