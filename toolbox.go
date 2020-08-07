package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
    
    "github.com/valyala/fasthttp"
    "github.com/fasthttp/router"

    "github.com/CloudyKit/jet"

    "github.com/bozso/gotoolbox/doc"
    //"github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/cli"
)

type TemplateServer struct {
    b doc.RenderBuilder
    errTpl, port string
    dev bool
}

func (ts *TemplateServer) SetCli(c *cli.Cli) {
    c.Var(&ts.b.Templates, "templates",
        "Paths to directory holding html templates")
    
    c.StringVar(&ts.errTpl, "errorTemplate", "",
        "Html template file path inside template directory for reporting errors.")
    
    c.StringVar(&ts.port, "port", "8080", "Http port to use.")
    c.BoolVar(&ts.dev, "dev", false,
        "Developement mode. If set templates are always reaload and not cached.")    
}

func (ts TemplateServer) Run() (err error) {
    r := ts.b.Build()
    
    r.Views.SetDevelopmentMode(ts.dev)
    //r.Views.SetAbortTemplateOnError(false)
            
    errView, err := r.Views.GetTemplate(ts.errTpl)
    if err != nil {
        return
    }
    
    errH := doc.NewErrorHandler(errView)
    
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

type TemplateRender struct {
    views *jet.Set
}

func (t TemplateRender) Error(w http.ResponseWriter, err error) {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    
}

func (t TemplateRender) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    view, err := t.views.GetTemplate(r.URL.Path)
    
    if err != nil {
        t.Error(w, err)
        return
    }
    
    err = view.Execute(w, nil, nil)

    if err != nil {
        t.Error(w, err)
        return
    }
}

func main() {
    c := cli.New("toolbox", "Useful functions.")
    
    c.AddAction("jet-server", "Render jet templates", &TemplateServer{})
    
    err := c.Run()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
}
