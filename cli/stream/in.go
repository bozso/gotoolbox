package stream

import (
    "os"
    "io"
    "bufio"

    "github.com/bozso/gotoolbox/path"
)

const stdinName Name = "stdin"

type InFile struct {
    stream
    io.ReadCloser
}

func (in *InFile) Set(s string) (err error) {
    f := path.New(s).ToFile()
    vf, err := f.ToValid()
    
    if err != nil {
        return
    }
    
    file, err := vf.Open()
    if err != nil {
        return
    }
    
    in.name, in.ReadCloser = Name(s), file
    return
}

func (in InFile) Scanner() (s *bufio.Scanner) {
    s = bufio.NewScanner(in.ReadCloser)
    return
}

func (in InFile) Reader() (r *bufio.Reader) {
    r = bufio.NewReader(in.ReadCloser)
    return
}

type In struct {
    InFile
}

func (in *In) Set(s string) (err error) {
    if l := len(s); l == 0 {
        in.name, in.ReadCloser = stdinName, os.Stdin
        return
    }
    
    err = in.InFile.Set(s)
    return
}
