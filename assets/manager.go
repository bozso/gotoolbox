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
	URL    string `json:"url"`
	MoveTo string `json:"move_to,omitempty"`
}

func (m Manager) download(rawURL string) (down path.Path, err error) {
	url, err := ParseUrl(rawURL)
	if err != nil {
		return
	}

	down, err = Download(m.download, url)
	if err != nil {
		return
	}

	return
}

func (m Manager) Download(t Task) (down path.ValidFile, err error) {
	const op = Operation("Manager.Download")

	m.download(t.Url)
	if err != nil {
		err = op.Error(err)
	}

	return
}

type ParallelManager struct {
	manager Manager
}
