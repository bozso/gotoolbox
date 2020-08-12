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
    Builder RenderBuilder
    errTpl string
    urlPrefix string
    dev bool
}

func (ts *TemplateServer) SetCli(c *cli.Cli) {
    ts.port.SetCli(c)
    
    c.Var(&ts.Builder.Templates, "templates",
        "Paths to directory holding html templates")
    
    c.StringVar(&ts.errTpl, "errorTemplate", "",
        "Html template file path inside template directory for reporting errors")

    c.StringVar(&ts.urlPrefix, "renderPrefix", "render",
        "Root path to use for template rendering")
    
    c.BoolVar(&ts.dev, "dev", false,
        "Developement mode. If set templates are always reaload and not cached")    
}

func (ts TemplateServer) Setup() (r Render, rout *router.Router, err error) {
    r = ts.Builder.Build()
    
    r.Views.SetDevelopmentMode(ts.dev)
    //r.Views.SetAbortTemplateOnError(false)
    
    errView, err := r.Views.GetTemplate(ts.errTpl)
    if err != nil {
        return
    }
    
    errH := NewErrorTemplate(errView)
    
    rout = router.New()
    rout.GET(fmt.Sprintf("/%s/{filepath:*}", ts.urlPrefix),
        errH.NewHandler(&r).Handle)
    
    fs := &fasthttp.FS{
        Root: "/",
        Compress: true,
    }
    
    rout.GET("/{path:*}", fs.NewRequestHandler())
    
    return
}

func (ts TemplateServer) Run() (err error) {
    _, rout, err := ts.Setup()
    
    if err != nil {
        return err
    }
    
    address := fmt.Sprintf(":%s", ts.port)
    log.Printf("Server starting on adrress: %s", address)
    err = fasthttp.ListenAndServe(address, rout.Handler)

    return
}

type RenderBuilder struct {
    Templates path.Dir
    funcs map[string]jet.Func
}

func NewRenderBuilder() (r RenderBuilder) {
    r.funcs = make(map[string]jet.Func)
    return
}

func (r *RenderBuilder) AddFunc(name string, f jet.Func) {
    r.funcs[name] = f
}

func (r RenderBuilder) Build() (rr Render) {
    v := jet.NewHTMLSet(r.Templates.GetPath())
    
    if funcs := r.funcs; funcs != nil {
        for key, val := range funcs {
            v.AddGlobalFunc(key, val)
        }
    }
    
    return Render{
        Views: v,
        d: New(),
        pool: bpool.NewBufferPool(defaultBufferPoolSize),
    }
}

type Handler interface {
    Handle(*fasthttp.RequestCtx) error
}
