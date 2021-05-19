package doc

import (
	"github.com/bozso/gotoolbox/cli"
	"github.com/bozso/gotoolbox/cli/stream"
)

type TemplateRender struct {
	builder  RenderBuilder
	template string
	out      stream.Out
}

func (t *TemplateRender) Default() {
	t.builder.Default()
	t.out.Default()
}

func (t *TemplateRender) SetCli(c *cli.Cli) {
	t.builder.SetCli(c)
	c.StringVar(&t.template, "template", "", "template file to render")
	c.Var(&t.out, "out", "output stream")
}

func (t TemplateRender) Run() (err error) {
	r := t.builder.Build()

	tpl, err := r.Views.GetTemplate(t.template)
	if err != nil {
		return err
	}

	err = tpl.Execute(t.out, nil, r.d)
	return
}
