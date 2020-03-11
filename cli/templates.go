package cli

import (
    "html/template"

    "github.com/bozso/gotoolbox/cli/stream"
)

type Templates struct {
    files ZeroOrMore
    name string
    out stream.Out
    jsonFile stream.InFile
    jsonOpt Optional
}

type TplName string

func (t *Templates) SetCli(c *Cli) {
    c.Var(&t.files, "templatePaths",
        "Comma separeted path to template files.")
    c.Var(&t.out, "out",
        "Output destination. By default writes to stdout.")
    
    c.StringVar(&t.name, "name", "", "Name of new template instance.")
    
    t.jsonOpt.SetCli(c, &t.jsonFile, "json", "Optional json file.")
}

func (t Templates) Run() (err error) {
    paths, err := t.files.ToPaths()
    if err != nil {
        return
    }
    
    tpl := template.New(t.name)
    
    for _, p := range paths {
        f, Err := p.ToFile()
        if err != nil {
            err = Err
            return
        }
        
        b, Err := f.ReadAll()
        if err != nil {
            err = Err
            return
        }
        
        _, Err = tpl.Parse(string(b))
        if err != nil {
            err = Err
            return
        }
    }
    
    
    var ctx map[string]interface{}
    
    if t.jsonOpt.IsSet() {
        ctx = make(map[string]interface{})
        
        // TODO: read jsonfile
    }
    
    
    w := t.out.BufWriter()
    
    err = tpl.Execute(w, ctx)
    
    return
}

