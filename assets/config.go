package assets

import (
	"io/fs"
)

type ConfigPayload struct {
	Cache  string `json:"cache"`
	Output string `json:"output"`
}

type Config struct {
	Cache  fs.FS
	Output fs.FS
}
