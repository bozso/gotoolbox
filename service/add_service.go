package service

import (
    "github.com/gofiber/fiber"
)

type ServiceAdder struct {
    rewriter PathRewriter
    router Router
}

func NewServiceAdder(router Router) (s ServiceAdder) {
    return ServiceAdder {
        router: router,
        rewriter: noRewriter,
    }
}

func (s ServiceAdder) WithRewriter(rewriter PathRewriter) (ss ServiceAdder) {
    s.rewriter = rewriter
    return s
}

func (s ServiceAdder) WithPath(path string) (ss ServiceAdder) {
    s.rewriter = SubPath(path)
    return s
}

func (s *ServiceAdder) Add(method Method, path string, handler fiber.Handler) {
    s.router.Add(string(method), s.rewriter.RewritePath(path), handler)
}

func (s *ServiceAdder) Delete(path string, handler fiber.Handler) {
    s.Add(Delete, path, handler)
}

func (s *ServiceAdder) Get(path string, handler fiber.Handler) {
    s.Add(Get, path, handler)
}

func (s *ServiceAdder) Post(path string, handler fiber.Handler) {
    s.Add(Post, path, handler)
}

func (s *ServiceAdder) Put(path string, handler fiber.Handler) {
    s.Add(Put, path, handler)
}

