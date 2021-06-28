package assets

import (
	"net/url"

	"github.com/bozso/gotoolbox/extract"
)

type Manager struct {
	download  Downloader
	extractor extract.Extractor
	config    Config
}

type Task struct {
	Url    string `json:"url"`
	MoveTo string `json:"move_to,omitempty"`
}

func (m Manager) Download(t Task) (down path.ValidFile, err error) {
	const op = Operation("Manager.Download")

	url, err := ParseUrl(t.Url)
	if err != nil {
		return op.Error(err)
	}

	down, err = Download(m.download, url)
	if err != nil {
		return op.Error(err)
	}

	return
}

type ParallelManager struct {
	manager Manager
}
