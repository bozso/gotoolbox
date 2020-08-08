package doc

import (
    "github.com/valyala/fasthttp"
    "github.com/CloudyKit/jet"
    "github.com/oxtoacart/bpool"
)

const defaultBufferPoolSize = 16

type Render struct {
    d Doc
    Views *jet.Set
    pool *bpool.BufferPool
}

func (r RenderBuilder) Build() (rr Render) {
    return Render{
        Views: jet.NewHTMLSet(r.Templates.GetPath()),
        d: New(),
        pool: bpool.NewBufferPool(defaultBufferPoolSize),
    }
}

func (r Render) Handle(ctx *fasthttp.RequestCtx) (err error) {
    path := ctx.UserValue("filepath").(string)
    view, err := r.Views.GetTemplate(path)

    if err != nil {
        return
    }
    
    buf := r.pool.Get()
    defer r.pool.Put(buf)
    
    err = view.Execute(buf, nil, r.d)
    if err != nil {
        return
    }

    ctx.SetContentType("text/html")
    _, err = buf.WriteTo(ctx.Response.BodyWriter())
    return
}
