package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/bozso/gotoolbox/cli"
	"github.com/bozso/gotoolbox/path"
)

type JetConfig struct {
	Dirs []string               `json:"dirs,omitempty"`
	Vars map[string]interface{} `json:"vars"`
}

var defaultDirs = []string{"."}

type Jet struct {
	configFile path.File
	out        string
	template   string
}

func (j *Jet) SetCli(c *cli.Cli) {
	c.NewFlag().
		Name("config").
		Usage("configuration file to use").
		Var(&j.configFile)

	c.NewFlag().
		Name("template").
		Usage("name of the template file to be rendered").
		StringVar(&j.template, "")

	c.NewFlag().
		Name("out").
		Usage("where to render the output").
		StringVar(&j.out, "")
}

func (j Jet) Run() (err error) {
	var cfg JetConfig

	exists, err := j.configFile.Exists()
	if err != nil {
		return
	}

	if exists {
		f, err := os.Open(j.configFile.String())
		if err != nil {
			return err
		}

		err = json.NewDecoder(bufio.NewReader(f)).Decode(&cfg)
		if err != nil {
			return err
		}
	}

	dirs := defaultDirs
	if len(cfg.Dirs) > 0 {
		dirs = cfg.Dirs
	}

	set := jet.NewHTMLSet(dirs...)

	if err != nil {
		return err
	}

	set.AddGlobalFunc("fmt", func(a jet.Arguments) (v reflect.Value) {
		a.RequireNumOfArguments("fmt", 1, -1)
		l := a.NumOfArguments()
		args := make([]interface{}, l-1)

		for ii := 1; ii < l; ii++ {
			args[ii-1] = a.Get(ii).Interface()
		}

		return reflect.ValueOf(fmt.Sprintf(a.Get(0).String(), args...))
	})

	tpl, err := set.GetTemplate(j.template)
	if err != nil {
		return
	}

	switch {
	case len(j.out) == 0 || strings.ToLower(j.out) == "stdout":
		out := os.Stdout
		err = tpl.Execute(out, nil, cfg.Vars)

		return err
	default:
		f, err := os.Create(j.out)
		if err != nil {
			return err
		}

		out := bufio.NewWriter(f)
		err = tpl.Execute(out, nil, cfg.Vars)

		out.Flush()
		f.Close()

		return err
	}
}
