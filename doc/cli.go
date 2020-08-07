package doc

import (
    "fmt"
    "log"

    "github.com/valyala/fasthttp"
    "github.com/fasthttp/router"

    "github.com/CloudyKit/jet"

    "github.com/oxtoacart/bpool"
    
    "github.com/bozso/gotoolbox/cli"
    "github.com/bozso/gotoolbox/path"
)

type Port struct {
    s int
}

func (p *Port) SetCli(c *cli.Cli) {
    c.IntVar(&p.s, "port", 8080, "Http port to use.")
}

func (p Port) Prepend(s string) (address string) {
    return fmt.Sprintf("%s:%d", s, p.s)
}

func (p Port) Localhost() (address string) {
    return fmt.Sprintf(":%d", p.s)
}

type TemplateServer struct {
    port Port
    builder RenderBuilder
    errTpl string
    dev bool
}

func (ts *TemplateServer) SetCli(c *cli.Cli) {
    ts.port.SetCli(c)
    
    c.Var(&ts.builder.Templates, "templates",
        "Paths to directory holding html templates")
    
    c.StringVar(&ts.errTpl, "errorTemplate", "",
        "Html template file path inside template directory for reporting errors.")
    
    c.BoolVar(&ts.dev, "dev", false,
        "Developement mode. If set templates are always reaload and not cached.")    
}

func (ts TemplateServer) Run() (err error) {
    r := ts.builder.Build()
    
    r.Views.SetDevelopmentMode(ts.dev)
    //r.Views.SetAbortTemplateOnError(false)
            
    errView, err := r.Views.GetTemplate(ts.errTpl)
    if err != nil {
        return
    }
    
    errH := NewErrorHandler(errView)
    
    router := router.New()
    router.GET("/render/{filepath:*}", errH.New(r).Handle)
    
    fs := &fasthttp.FS{
        Root: "/",
        Compress: true,
        //PathRewrite: fasthttp.NewPathSlashesStripper(1),
    }
    
    router.GET("/{path:*}", fs.NewRequestHandler())
    
    address := fmt.Sprintf(":%s", ts.port)
    log.Printf("Server starting on adrress: %s", address)
    err = fasthttp.ListenAndServe(address, router.Handler)

    return
}

type RenderBuilder struct {
    Templates path.Dir
}

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

type Handler interface {
    Handle(*fasthttp.RequestCtx) error
}

type ErrorHandler interface {
    Handle(ctx *fasthttp.RequestCtx)
}

type ErrorHandlerFactory struct {
    template *jet.Template
}

func NewErrorHandler(template *jet.Template) (e ErrorHandlerFactory) {
    e.template = template
    return
}

func (e ErrorHandlerFactory) New(h Handler) (et ErrorTemplate) {
    et.template, et.handler = e.template, h
    return
}

type ErrorTemplate struct {
    handler Handler
    template *jet.Template
}

func (e ErrorTemplate) Handle(ctx *fasthttp.RequestCtx) {
    const format = "Error while executing error template: %s\n"
    
    err := e.handler.Handle(ctx)
    if err == nil {
        return
    }
    
    log.Printf("Error: %s\n", err)
    
    ctx.SetContentType("text/html")
    err = e.template.Execute(ctx.Response.BodyWriter(), nil, err)
    
    if err != nil {
        ctx.Response.AppendBodyString(fmt.Sprintf(format, err))
    }
}
