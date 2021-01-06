package path

import (
    "io"
)

type ReadCloserCreator interface {
    CreateReadCloser() (r io.ReadCloser, err error)
}

type ReadResource struct {
    creator ReadCloserCreator
}

func (r ReadResource) Use(fn func (io.ReadCloser) error) (err error) {
    rc, err := r.creator.CreateReadCloser()
    if err != nil {
        return
    }

    if err = fn(rc); err != nil {
        return
    }
    return rc.Close()
}
