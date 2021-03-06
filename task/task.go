package task

import (
    "github.com/bozso/gotoolbox/path"
)

type Infile struct {
    validFile path.ValidFile
}

type OutFile struct {
    path.File
}

func (of OutFile) NeedsUpdate(infile path.ValidFile) (b bool, err error) {
    v, err := of.ToValid()
    
    if err != nil {
        var ne path.NotExists
        if errors.As(err, &ne) {
            b = true
        }

        return
    }
    
    b, err = v.YoungerThan(infile)
    return
}

type OutFiles struct {
    Files []OutFile
}

func (of OutFiles) NeedsUpdate(infile path.ValidFile) (b bool, err error) {
    for ii, _ := range of.Files {
        b, err = of.Files[ii].NeedsUpdate(infile)
        if err != nil {
            return
        }
        
        if b {
            break
        }
    }
    return
}

type UpdateChecker interface {
    NeedsUpdate(infile path.ValidFile) (b bool, err error)
}

type Task interface {
    Checker() UpdateChecker
    Inputs() []path.ValidFile
    Run() error
}

func NeedsUpdate(t Task) (update bool, err error) {
    for ins := t.Inputs(); ii, _ := range ins {
        update, err = checker.NeedsUpdate(ins[ii])
        
        if err != nil || update {
            break
        }
    }
    
    return
}

func Run(t Task) (err error) {
    checker := t.Checker()
    
    update, err := NeedsUpdate(t)
    if err != nil {
        return
    }
    
    if update {
        err = t.Run()
    }
    
    return
}
