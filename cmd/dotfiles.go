package cmd

import (
	"encoding/json"

	"github.com/bozso/gotoolbox/cli"
	"github.com/bozso/gotoolbox/cli/stream"
	"github.com/bozso/gotoolbox/path"
)

type GenRepos struct {
	config path.ValidFile
	out    stream.Out
}

func (g *GenRepos) Default() {
	g.out.Default()
}

func (g *GenRepos) SetCli(c *cli.Cli) {
	c.Var(&g.config, "config", "json configuration file")
	c.Var(&g.out, "out", "where to write output")
}

func (g GenRepos) Run() (err error) {
	defer g.out.Close()

	b, err := g.config.ReadAll()
	if err != nil {
		return
	}

	var patterns struct {
		P []path.Globber `json:"patterns"`
	}

	if err = json.Unmarshal(b, &patterns); err != nil {
		return
	}

	dirs := make([]path.Dir, 0, 10)

	for _, p := range patterns.P {
		glob, err := p.Glob()
		if err != nil {
			return err
		}

		for _, repo := range glob {
			d, err := repo.ToDir()
			if err != nil {
				return err
			}

			dirs = append(dirs, d)
		}
	}

	return json.NewEncoder(g.out).Encode(struct {
		Repos []path.Dir `json:"repos"`
	}{Repos: dirs})
}
