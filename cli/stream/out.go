package stream

import (
    "os"
    "io"
    "bufio"

    "github.com/bozso/gotoolbox/path"
)

const stdoutName Name = "stdout"

type OutFile struct {
    stream
    io.WriteCloser
}

func (o *OutFile) Set(s string) (err error) {
    p, err := path.New(s).ToFile()
    if err != nil {
        return
    }
    
    file, err := p.Open()
    if err != nil {
        return
    }
    
    o.name, o.WriteCloser = Name(s), file
    return
}

func (o Out) BufWriter() (w *bufio.Writer) {
    w = bufio.NewWriter(o.WriteCloser)
    return
}

type Out struct {
    OutFile
}

func (o *Out) Set(s string) (err error) {
    if l := len(s); l == 0 {
        o.name, o.WriteCloser = stdoutName, os.Stdout
        return
    }
    
    err = o.OutFile.Set(s)
    return
}
