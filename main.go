package main

import (
    "fmt"
    "os"
    
    "github.com/bozso/gotoolbox/doc"
    "github.com/bozso/gotoolbox/cli"
)

func main() {
    c := cli.New("toolbox", "Useful functions.")
    
    c.AddAction("jet-server", "Render jet templates",
        &doc.TemplateServer{})
    
    err := c.Run()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
}
