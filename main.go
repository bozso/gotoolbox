package main

import (
    "fmt"
    "os"
    "net/http"
    
    "github.com/gorilla/mux"
    "github.com/gorilla/rpc/v2"
    "github.com/gorilla/rpc/v2/json2"

    "github.com/bozso/gotoolbox/doc"
    "github.com/bozso/gotoolbox/cli"
    "github.com/bozso/gotoolbox/services"
)


type Service struct {
    port doc.Port
}

func (s *Service) SetCli(c *cli.Cli) {
    s.port.SetCli(c)
    return
}

func (sv Service) Run() (err error) {
    s := rpc.NewServer()
    s.RegisterCodec(json2.NewCodec(), "application/json")
    s.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")

    err = s.RegisterService(services.EncoderService{}, "")

    r := mux.NewRouter()
    r.Handle("/rpc", s)

    
    return http.ListenAndServe(sv.port.Localhost(), r)
}

func main() {
    c := cli.New("toolbox", "Useful functions.")
    
    c.AddAction("jet-server", "render jet templates through a web server",
        &doc.TemplateServer{})

    c.AddAction("template", "render jet templates",
        &doc.TemplateRender{})
    
    c.AddAction("service", "start services",
        &Service{})

    err := c.Run()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
}
