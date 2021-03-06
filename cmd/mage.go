package cmd

import (
    "os"
    "fmt"

    "github.com/bozso/gotoolbox/cli"
    "github.com/magefile/mage/mage"
)

type Mage struct {}

func (m *Mage) SetCli(c *cli.Cli) {}

func (m Mage) Run() (err error) {
    args := os.Args

    if l := len(args); l < 2 {
        return fmt.Errorf("expected at least two arguments")
    }

    os.Exit(mage.ParseAndRun(os.Stdout, os.Stderr, os.Stdin, args[2:]))
    return
}
