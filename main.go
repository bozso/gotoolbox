package main

import (
    "fmt"
    "os"
    "strings"
    "net/http"
    
    "github.com/gorilla/mux"
    "github.com/gorilla/rpc/v2"
    "github.com/gorilla/rpc/v2/json2"

    tth "github.com/buildkite/terminal-to-html"

    "github.com/bozso/gotoolbox/doc"
    "github.com/bozso/gotoolbox/cli"
    "github.com/bozso/gotoolbox/command"
    "github.com/bozso/gotoolbox/cli/stream"
    "github.com/bozso/gotoolbox/services"
    "github.com/bozso/gotoolbox/repository"
)

type Repositories struct {
    config repository.Config
    vcs string
    command string
    intoHtml bool
    out stream.Out
}

func (r *Repositories) SetCli(c *cli.Cli) {
    c.Var(&r.config, "config", "json configuration of repository list")
    c.BoolVar(&r.intoHtml, "html", false,
        "whether to convert output to html")
    c.StringVar(&r.vcs, "vcs", "git",
        "type of the version control system to use")
    c.StringVar(&r.command, "command", "status",
        "which type of command to use")
    
    c.Var(&r.out, "out", "where to write the output")
}

func (r Repositories) Run() (err error) {
    var vcs command.VCS
    
    switch strings.ToLower(r.vcs) {
    case "git":
        vcs = command.DefaultGit()
    default:
        err = fmt.Errorf("unknown version control system '%s'", r.vcs)
        return
    }
    
    m := r.config.IntoManager(vcs)
    
    //fmt.Printf("%#v\n", m)
    
    var out []byte
    switch strings.ToLower(r.command) {
    case "status":
        out, err = m.Status()
    default:
        err = fmt.Errorf("unknown manager command '%s'", r.command)
    }
    
    if err != nil {
        return
    }
    
    if r.intoHtml {
        out = tth.Render(out)
    }
    
    _, err = r.out.Write(out)
    return
}

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
    
    c.AddAction("repositories",
        "manage version control system repositories",
        &Repositories{})

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
