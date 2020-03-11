package cli

import (
    "fmt"
    "flag"

    "github.com/bozso/gotoolbox/errors"
)

type (
    Action interface {
        SetCli(*Cli)
        Run() error
    }
    
    subcommand struct {
        action Action
        c Cli
    }
    
    subcommands map[string]subcommand
    
    Cli struct {
        desc string
        *flag.FlagSet
        subcommands
    }
)

func (c subcommands) Keys() []string {
    s := make([]string, len(c))
    
    ii := 0
    for k := range c {
        s[ii] = k
        ii++
    }
    return s
}

func New(name, desc string) (c Cli) {
    c.desc = desc
    c.FlagSet = flag.NewFlagSet(name, flag.ContinueOnError)
    c.subcommands = make(map[string]subcommand)
    
    return c
}

func (c *Cli) AddAction(name, desc string, act Action) {
    c.subcommands[name] = subcommand{
        action: act,
        c: New(name, desc),
    }
}

func (c Cli) HasSubcommands() bool {
    return c.subcommands != nil && len(c.subcommands) != 0
}

func (c Cli) Usage() {
    fmt.Printf("Program: %s. Description: %s\n",
        Bold.Wrap(c.Name()), c.desc)
    c.PrintDefaults()
    
    if c.HasSubcommands() {
        fmt.Printf("\nAvailable subcommands: %s\n", c.subcommands.Keys())
    }
}

func (c Cli) Run(args []string) (err error) {
    //ferr := merr.Make("Cli.Run")
    
    if !c.HasSubcommands() {
        err = c.Parse(args)
        return
    }
    
    l := len(args)
    
    if l < 1 {
        fmt.Printf("Expected at least one parameter specifying subcommand.\n")
        c.Usage()
        return nil
    }
    
    // TODO: check if args is long enough
    mode := args[0]
    
    if mode == "-help" || mode == "-h" {
        c.Usage()
        return nil
    }
    
    subcom, ok := c.subcommands[mode]
    
    if !ok {
        return errors.UnrecognizedMode(mode, c.Name())
    }
    
    subc, act := &subcom.c, subcom.action
    subcom.action.SetCli(subc)
    
    err = subc.Parse(args[1:])
    
    if err != nil {
        return
    }
    
    return act.Run()
}

