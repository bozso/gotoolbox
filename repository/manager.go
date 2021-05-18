package repository

import (
	"bytes"
	"log"

	"github.com/bozso/gotoolbox/command"
	"github.com/bozso/gotoolbox/path"
)

type Manager struct {
	Repositories
	Vcs command.VCS
}

func NewManager(vcs command.VCS) (m Manager) {
	return Manager{
		Vcs:          vcs,
		Repositories: make(Repositories, 0, 10),
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
		dir := m.Repositories[ii].Directory
		err = dir.Chdir()
		if err != nil {
			return
		}

		b, err = m.Vcs.Status()
		if err != nil {
			log.Printf("error while processing directory '%s': %s\n",
				dir, err)
			continue
		}
		buf.Write(b)
	}

	b, err = buf.Bytes(), curr.Chdir()
	return
}
