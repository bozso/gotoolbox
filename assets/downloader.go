package assets

import (
	"net/url"

	"github.com/bozso/gotoolbox/path"
)

type Downloader interface {
	Download(*url.URL) (path.ValidFile, error)
}

func ParseUrl(raw string) (url *url.URL, err error) {
	const op = Operation("parsing url")

	url, err = url.Parse(raw)
	if err != nil {
		err = op.Error(err)
	}
	return
}

func Download(d Downloader, url *url.URL) (p path.ValidFile, err error) {
	const op = Operation("downloading asset")

	p, err = d.Download(url)
	if err != nil {
		err = DownloadError{
			URL: url,
			err: err,
		}
	}

	return

}
