package service

import (
    "fmt"

    "github.com/gofiber/fiber"
)

type Router interface {
    Add(method string, path string, handler fiber.Handler)
}

type PathRewriter interface {
    RewritePath(in string) (out string)
}

type noRewrite struct {}

func (_ noRewrite) RewritePath(in string) (out string) {
    return in
}

var noRewriter = noRewrite{}

type WithPath struct {
    subpath string
}

func SubPath(path string) (w WithPath) {
    return WithPath {
        subpath: path,
    }
}

func (w WithPath) RewritePath(in string) (out string) {
    return fmt.Sprintf("%s/%s", w.subpath, in)
}
