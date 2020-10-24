package repository

import (
    "fmt"
    "encoding/json"

    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/command"
)

type Config struct {
    Repositories `json:"repos"`
}

func (c Config) String() (s string) {
    return fmt.Sprintf("repositories: %v", c.Repositories)
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
    
    return json.Unmarshal(b, &c)
}

func (c Config) IntoManager(vcs command.VCS) (m Manager) {
    return Manager{
        Repositories: c.Repositories,
        Vcs: vcs,
    }
}
