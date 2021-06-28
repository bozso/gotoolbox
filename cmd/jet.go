package cmd

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/bozso/gotoolbox/cli"
	"github.com/bozso/gotoolbox/path"
)

type JetConfig struct {
	Dirs []string
}

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
	cfg := DefaultJetConfig

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

	set := jet.NewHTMLSet(cfg.Dirs...)

	if err != nil {
		return err
	}

	tpl, err := set.GetTemplate(j.template)
	if err != nil {
		return
	}

	var out io.Writer

	switch {
	case len(j.out) == 0 || strings.ToLower(j.out) == "stdout":
		out = os.Stdout
	default:
		f, err := os.Open(j.out)
		if err != nil {
			return err
		}

		out = bufio.NewWriter(f)
		defer f.Close()
	}

	return tpl.Execute(out, nil, nil)
}

var DefaultJetConfig = JetConfig{
	Dirs: []string{"."},
}
