package main

import (
	"os"
	"os/exec"

	"github.com/bozso/gotoolbox/cli"
	"github.com/bozso/gotoolbox/path"
)

const taskCmd = "task"

var (
	parentDir = path.New("..")
)

type Tasker struct {}

func (t *Tasker) SetCli(c *cli.Cli) {}

func (t Tasker) Run() (err error) {
	taskFile := path.New("Taskfile.yml")

	for {
		b, err := taskFile.Exists()

		if err != nil {
			return err
		}

		if b {
			break
		} else {
			taskFile = path.New(parentDir.Join(taskFile.String()).String())
		}
	}

	cmd := exec.Command(taskCmd,
		append(os.Args[2:], "--taskfile", taskFile.String())...)

	out, err := cmd.CombinedOutput()
	println(string(out))
	return
}
