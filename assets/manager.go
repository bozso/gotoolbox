package assets

import (
	"github.com/bozso/gotoolbox/extract"
)

type Manager struct {
	download  Downloader
	extractor extract.Extractor
	config    Config
}
