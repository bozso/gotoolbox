package optional

import (
    "flag"

    "github.com/bozso/gotoolbox/cli"
)

type Optional struct {
    flag.Value
    set bool
}

func (o Optional) IsSet() bool {
    return o.set
}

func (o *Optional) SetCli(c *cli.Cli, val flag.Value, name, descr string) {
    o.set = false
    c.Var(o.Setup(val), name, descr)
    
}

func (o *Optional) Setup(val flag.Value) *Optional {
    o.Value = val
    return o
}

func (o *Optional) Set(s string) (err error) {
    if len(s) == 0 {
        err = nil
        o.set = false
    } else {
        err = o.Value.Set(s)
        
    }
    
    if err != nil {
        o.set = true
    }
    
    return
}
