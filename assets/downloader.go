package assets

import (
	"net/url"

	"github.com/bozso/gotoolbox/path"
)

type Downloader interface {
	download(uri url.URL, out path.Path) error
}
