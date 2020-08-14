package main

import (
    "fmt"
    "os"
    
    "github.com/bozso/gotoolbox/doc"
    "github.com/bozso/gotoolbox/cli"
)

func main() {
    c := cli.New("toolbox", "Useful functions.")
    
    c.AddAction("jet-server", "render jet templates through a web server",
        &doc.TemplateServer{})

    c.AddAction("template", "render jet templates",
        &doc.TemplateRender{})
    
    err := c.Run()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
}
