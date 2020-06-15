package task

import (
    "github.com/bozso/gotoolbox/path"
)

type Meta struct {
    Infiles []path.ValidFile
    Outfiles []path.File    
}

type Task interface {
    GetMeta() Meta
    Run() error
}

func NeedsToRun(t Task) (b bool, err error) {
    b = true
    m := t.GetMeta()
    
    for ii, _ := range m.Outfiles {
        exists, err := m.Outfiles[ii].Exist()
        if err != nil {
            return b, err
        }
        
        if !exists {
            return true, nil
        }
    }
    
    for ii, _ := range m.Infiles {
        
    } 
}
