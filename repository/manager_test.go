package repository

import (
    "log"
    "testing"

    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/command"
)

func testStatus() (err error) {
    d, err := path.New(".").ToDir()
    if err != nil {
        return
    }
    
    r := New(d)
    
    m := NewManager(command.NewGit(command.NewExecutable("git")))
    m.AddRepo(r)
    
    -, err := m.Status()
    return
}

func TestStatus(t *testing.T) {
    if err := testStatus(); err != nil {
        t.Fatalf("error: %s\n", err)
    }
}
