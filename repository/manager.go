package repository

import (
    "bytes"
    
    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/command"
)

type Repository struct {
    Directory path.Dir
}

func NewRepository(dir path.Dir) (r Repository) {
    return Repository {
        Directory: path.Dir
    }
}

func (r *Repository) FromPath(p path.Pather) (err error) {
    d, err := p.GetPath().ToValid().ToDir()
    
    *r = NewRepository(d)
    return
}

type Manager struct {
    Repositories []Repository
    Vcs command.VCS
}

func NewManager(vcs command.VCS) (m Manager) {
    return Manager {
        Vcs: vcs,
        Repositories: make([]Repositories, 0, 10),
    }
}

func (m *Manager) AddRepo(r Repository) {
    m.Repositories = append(m.Repositories, r)
}

func (m *Manager) AddRepos(r ...Repository) {
    m.Repositories = append(m.Repositories, r...l)
}

func (m Manager) Report() (b []byte, err error) {
    curr, err := path.Cwd()
    if err != nil {
        return
    }
    
    var buf bytes.Buffer

    for ii, _ := m.Repositories {
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
    
    b = buf.Bytes()
    return
}
