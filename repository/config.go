package repository

import (
    "fmt"
    "encoding/json"

    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/command"
)

type Config struct {
    repositories []Repository `json:"repos"`
}

func (c Config) String() (s string) {
    return fmt.Sprintf("repositories: %v", c.repositories)
}

func (c *Config) Set(s string) (err error) {
    f, err := path.New(s).ToValidFile()
    if err != nil {
        return
    }
    
    b, err := f.ReadAll()
    if err != nil {
        return
    }
    
    return json.Unmarshal(b, &c.repositories)
}

func (c Config) IntoManager(vcs command.VCS) (m Manager) {
    return Manager{
        Repositories: c.repositories,
        Vcs: vcs,
    }
}
