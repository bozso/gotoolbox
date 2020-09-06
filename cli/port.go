package cli

import (
    "fmt"
)

type Port struct {
    s int
    mode string
}

func (p *Port) Default() {
    p.s = 8080
    p.mode = "localhost"
}

func (p *Port) SetCli(c *Cli) {
    c.IntVar(&p.s, "port", 8080, "http port to use")
    c.StringVar(&p.mode, "mode", "localhost", "mdoe to use")
}

func (p Port) Prepend(s string) (address string) {
    return fmt.Sprintf("%s:%d", s, p.s)
}

func (p Port) Localhost() (address string) {
    return fmt.Sprintf(":%d", p.s)
}

func (p Port) HostName() (s string) {
    switch mode := p.mode; mode {
    case "localhost":
        s = p.Localhost()
    default:
        s = p.Prepend(mode)
    }
    return
}
