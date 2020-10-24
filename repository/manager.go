package repository

import (
    "bytes"
    
    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/command"
)

type Manager struct {
    Repositories []Repository
    Vcs command.VCS
}

func NewManager(vcs command.VCS) (m Manager) {
    return Manager {
        Vcs: vcs,
        Repositories: make([]Repository, 0, 10),
    }
}

func (m *Manager) AddRepo(r Repository) {
    m.Repositories = append(m.Repositories, r)
}

func (m *Manager) AddRepos(r ...Repository) {
    m.Repositories = append(m.Repositories, r...)
}

func (m Manager) Status() (b []byte, err error) {
    curr, err := path.Cwd()
    if err != nil {
        return
    }
    
    var buf bytes.Buffer

    for ii, _ := range m.Repositories {
        err = m.Repositories[ii].Directory.Chdir()
        if err != nil {
            return
        }
        
        b, err = m.Vcs.Status()
        if err != nil {
            return
        }
        buf.Write(b)
    }
    
    b, err = buf.Bytes(), curr.Chdir()
    return
}
