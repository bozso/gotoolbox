package path

import (
    "fmt"
    "io"
    "bufio"
)

/*
ReadCreator represents an object that can create a ReadCloser.
*/
type ReaderCreator interface {
    CreateReader() (r io.ReadCloser, err error)
}

type BufferedCreator struct {
    creator ReaderCreator
}

func (b BufferedCreator) CreateReader() (reader io.ReadCloser, err error) {
    r, err := b.creator.CreateReader()
    if err != nil {
        return
    }

    reader = &BufferedReadCloser {
        ReadCloser: r,
        Reader: bufio.NewReader(r),
    }
    return
}

type BufferedReadCloser struct {
    io.ReadCloser
    *bufio.Reader
}

func (b *BufferedReadCloser) Read(p []byte) (n int, err error) {
    return b.Reader.Read(p)
}


/*
UseReader creates a ReadCloser object, applies the fn function to it
(uses the resource), and finally closes it. This ensures that the error
reported by the Close method is returned and not ignored and clearly
denotes the lifetime of a ReadCloser instance.
*/
func UseReader(r ReaderCreator, fn func (io.Reader) error) (err error) {
    rc, err := r.CreateReader()
    if err != nil {
        return
    }

    if err = fn(rc); err != nil {
        return closeWrap(err, rc)
    }
    return rc.Close()
}

/*
UseAsScanner creates a ReaderCloser and wraps it inside a bufio.Scanner
struct to be used by fn.
*/
func UseAsScanner(r ReaderCreator, fn func(*bufio.Scanner) error) (err error) {
    rc, err := r.CreateReader()
    if err != nil {
        return
    }

    scanner := bufio.NewScanner(rc)

    if err = fn(scanner); err != nil {
        return closeWrap(err, rc)
    }

    return rc.Close()
}

type WriterCreator interface {
    CreateWriter() (io.WriteCloser, error)
}

func UseWriter(w WriterCreator, fn func(io.Writer) error) (err error) {
    wc, err := w.CreateWriter()
    if err != nil {
        return
    }

    if err = fn(wc); err != nil {
        return closeWrap(err, wc)
    }
    return wc.Close()
}


type CloseError struct {
    err error
}

func (ce CloseError) Error() (s string) {
    return fmt.Sprintf(
        "failed to close resource while while handling error %s", ce.err)
}

func closeWrap(err error, c io.Closer) (Err error) {
    if err := c.Close(); err != nil {
        err = CloseError{ err: err }
    }

    return err
}
