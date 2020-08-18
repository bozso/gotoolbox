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

func DefaultOut() OutFile {
    return OutFile{
        stream: stream{name: stdoutName},
        WriteCloser: os.Stdout,
    }
}

func (o *OutFile) Set(s string) (err error) {
    f := path.New(s).ToFile()
    
    file, err := f.Create()
    if err != nil {
        return
    }
    
    o.name, o.WriteCloser = Name(s), file
    return
}

func (o OutFile) BufWriter() (w *bufio.Writer) {
    w = bufio.NewWriter(o.WriteCloser)
    return
}

type Out struct {
    OutFile
}

func (o *Out) Default() {
    o.name, o.WriteCloser = stdoutName, os.Stdout
}

func (o *Out) Set(s string) (err error) {
    if len(s) != 0 {
        err = o.OutFile.Set(s)
    }
    
    return
}
