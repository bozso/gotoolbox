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
    f := path.NewFile(s)
    
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

func (o *Out) Set(s string) (err error) {
    println("asd")
    println(s)
    if l := len(s); l == 0 {
        o.name, o.WriteCloser = stdoutName, os.Stdout
        return
    }
    
    err = o.OutFile.Set(s)
    return
}
