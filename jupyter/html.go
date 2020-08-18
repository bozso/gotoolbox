package jupyter

import (
    "sync"
    "fmt"

    "github.com/bozso/gotoolbox/errors"
)

type RenderHTML interface {
    ToHtml() (b []byte)
}

type Json map[string]interface{}

type RenderJson interface {
    ToJson() (m Json)
}

type ID int64

func (id ID) String() (s string) {
    return fmt.Sprintf("%d", int64(id))
}

type HtmlMap map[ID]RenderHTML
type JsonMap map[ID]RenderJson

type Entities struct {
    mutex sync.Mutex
    html HtmlMap
    json JsonMap
}

func NewEntities() (e Entities) {
    e.html, e.json = make(HtmlMap), make(JsonMap)
    return
}

func (e *Entities) SetHtml(id ID, h RenderHTML) {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    
    e.html[id] = h
}

func (e *Entities) SetJson(id ID, j RenderJson) {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    
    e.json[id] = j
}

func (e *Entities) GetHTML(id ID) (b []byte, err error) {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    
    h, ok := e.html[id]
    
    if !ok {
        err = errors.KeyNotFound(id.String())
        return
    }
    
    b = h.ToHtml()
    return
}

func (e *Entities) GetJson(id ID) (j Json, err error) {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    
    js, ok := e.json[id]
    
    if !ok {
        err = errors.KeyNotFound(id.String())
        return
    }
    
    j = js.ToJson()
    return
}
