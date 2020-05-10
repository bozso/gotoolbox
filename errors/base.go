package errors

type Base struct {
    err error
}

func (e *Base) Wrap(err error) (b *Base) {
    e.err = err
    return e
}

func (e Base) Unwrap() error {
    return e.err
}
