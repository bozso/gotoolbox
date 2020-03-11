package cli

import (
    "html/template"
)

type Templates struct {
    files ZeroOrMore
    name string
    
}

type TplName string

func (t *Templates) SetCli(c *Cli) {
    c.Var(&t.files, "templatePaths",
        "Comma separeted path to template files.")
    c.StringVar(&t.name, "name", "", "Name of new template instance.")
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
    
    tpl.Execute(
    
    return nil
}

