package assets

import (
	"net/url"

	"github.com/bozso/gotoolbox/path"
)

type Downloader interface {
	Download(*url.URL, string) error
}

func ParseUrl(raw string) (url *url.URL, err error) {
	const op = Operation("parsing url")

	url, err = url.Parse(raw)
	if err != nil {
		err = op.Error(err)
	}
	return
}

func Download(d Downloader, url *url.URL, path string) (err error) {
	const op = Operation("downloading asset")

	err = d.Download(url, path)
	if err != nil {
		err = DownloadError{
			URL: url,
			err: err,
		}
	}

	return
}
