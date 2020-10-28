package main

import (
    "os"
    "os/exec"

    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/cli"
)

const taskCmd = "task"

var (
    parentDir = path.New("..")
)

type Tasker struct {
    
}

func (t *Tasker) SetCli(c *cli.Cli) {
    
}

func (t Tasker) Run() (err error) {
    taskFile := path.New("Taskfile.yml")
    
    for {
        b, err := taskFile.Exist()
        
        if err != nil {
            return err
        }
        
        if b {
            break
        } else {
            taskFile = parentDir.Join(taskFile.String())
        }
    }
    
    cmd := exec.Command(taskCmd,
        append(os.Args[2:], "--taskfile", taskFile.String())...)
    
    out, err := cmd.CombinedOutput()
    println(string(out))
    return
}
