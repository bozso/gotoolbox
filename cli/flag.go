package cli

import (
	"flag"
)

type FlagDesc struct {
	name, usage string
}

func (f FlagDesc) Name(name string) (fd FlagDesc) {
	f.name = name
	return f
}

func (f FlagDesc) Usage(usage string) (fd FlagDesc) {
	f.usage = usage
	return f
}

func (f FlagDesc) Var(c *Cli, val flag.Value) {
	c.Var(val, f.name, f.usage)
}

func (f FlagDesc) Bool(c *Cli, val bool) (b *bool) {
	return c.Bool(f.name, val, f.usage)
}

func (f FlagDesc) BoolVar(c *Cli, p *bool, val bool) {
	c.BoolVar(p, f.name, val, f.usage)
}

func (f FlagDesc) Float64(c *Cli, val float64) (fl *float64) {
	return c.Float64(f.name, val, f.usage)
}

func (f FlagDesc) Float64Var(c *Cli, p *float64, val float64) {
	c.Float64Var(p, f.name, val, f.usage)
}

func (f FlagDesc) Int(c *Cli, val int) (i *int) {
	return c.Int(f.name, val, f.usage)
}

func (f FlagDesc) IntVar(c *Cli, p *int, val int) {
	c.IntVar(p, f.name, val, f.usage)
}

func (f FlagDesc) Int64(c *Cli, val int64) (i *int64) {
	return c.Int64(f.name, val, f.usage)
}

func (f FlagDesc) Int64Var(c *Cli, p *int64, val int64) {
	c.Int64Var(p, f.name, val, f.usage)
}

func (f FlagDesc) String(c *Cli, val string) (s *string) {
	return c.String(f.name, val, f.usage)
}

func (f FlagDesc) StringVar(c *Cli, p *string, val string) {
	c.StringVar(p, f.name, val, f.usage)
}

func (f FlagDesc) Uint64(c *Cli, val uint64) (i *uint64) {
	return c.Uint64(f.name, val, f.usage)
}

func (f FlagDesc) Uint64Var(c *Cli, p *uint64, val uint64) {
	c.Uint64Var(p, f.name, val, f.usage)
}

func (f FlagDesc) Uint(c *Cli, val uint) (i *uint) {
	return c.Uint(f.name, val, f.usage)
}

func (f FlagDesc) UintVar(c *Cli, p *uint, val uint) {
	c.UintVar(p, f.name, val, f.usage)
}

func (c *Cli) NewFlag() (f FlagDescWithCli) {
	f.c = c
	return
}

type FlagDescWithCli struct {
	FlagDesc
	c *Cli
}

func (f FlagDescWithCli) Name(name string) (fd FlagDescWithCli) {
	f.name = name
	return f
}

func (f FlagDescWithCli) Usage(usage string) (fd FlagDescWithCli) {
	f.usage = usage
	return f
}

func (f FlagDescWithCli) Var(val flag.Value) {
	f.c.Var(val, f.name, f.usage)
}

func (f FlagDescWithCli) Bool(val bool) (b *bool) {
	return f.c.Bool(f.name, val, f.usage)
}

func (f FlagDescWithCli) BoolVar(p *bool, val bool) {
	f.c.BoolVar(p, f.name, val, f.usage)
}

func (f FlagDescWithCli) Float64(val float64) (fl *float64) {
	return f.c.Float64(f.name, val, f.usage)
}

func (f FlagDescWithCli) Float64Var(p *float64, val float64) {
	f.c.Float64Var(p, f.name, val, f.usage)
}

func (f FlagDescWithCli) Int(val int) (i *int) {
	return f.c.Int(f.name, val, f.usage)
}

func (f FlagDescWithCli) IntVar(p *int, val int) {
	f.c.IntVar(p, f.name, val, f.usage)
}

func (f FlagDescWithCli) Int64(val int64) (i *int64) {
	return f.c.Int64(f.name, val, f.usage)
}

func (f FlagDescWithCli) Int64Var(p *int64, val int64) {
	f.c.Int64Var(p, f.name, val, f.usage)
}

func (f FlagDescWithCli) String(val string) (s *string) {
	return f.c.String(f.name, val, f.usage)
}

func (f FlagDescWithCli) StringVar(p *string, val string) {
	f.c.StringVar(p, f.name, val, f.usage)
}

func (f FlagDescWithCli) Uint64(val uint64) (i *uint64) {
	return f.c.Uint64(f.name, val, f.usage)
}

func (f FlagDescWithCli) Uint64Var(p *uint64, val uint64) {
	f.c.Uint64Var(p, f.name, val, f.usage)
}

func (f FlagDescWithCli) Uint(val uint) (i *uint) {
	return f.c.Uint(f.name, val, f.usage)
}

func (f FlagDescWithCli) UintVar(p *uint, val uint) {
	f.c.UintVar(p, f.name, val, f.usage)
}
