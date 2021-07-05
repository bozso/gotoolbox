package assets

import (
	"io/fs"
	"os"
)

type ConfigPayload struct {
	Cache  string `json:"cache"`
	Output string `json:"output"`
}

func (cp ConfigPayload) Config() (c Config, err error) {
	return Config{
		Cache:  os.DirFS(cp.Cache),
		Output: os.DirFS(cp.Output),
	}, nil
}

type Config struct {
	Cache  fs.FS
	Output fs.FS
}
