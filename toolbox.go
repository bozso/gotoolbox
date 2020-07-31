package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
    
    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/cli"
    "github.com/CloudyKit/jet"
)

type Jet struct {
    isDev bool
    port string
    templateDir path.Dir
}

func (j *Jet) SetCli(c *cli.Cli) {
    c.BoolVar(&j.isDev, "dev", false, "Developement mode")
    
    c.Var(&j.templateDir, "templates",
        "Path to directory holding templates")
    
    c.StringVar(&j.port, "port", "8080", "Localhost port to use")
}

func (j Jet) Run() (err error) {
    views := jet.NewHTMLSet(j.templateDir.GetPath())
    views.SetDevelopmentMode(j.isDev)
    
    tr := TemplateRender{views}
    
    http.Handle("/", http.FileServer(http.Dir("/")))
    http.Handle("/render/", http.StripPrefix("/render/", tr))

    address := fmt.Sprintf(":%s", j.port)
    
    log.Printf("Server starting at port: %s\n", address)
    err = http.ListenAndServe(address, nil)
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
    
    c.AddAction("jet", "Redner jet templates", &Jet{})
    
    err := c.Run()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
}
