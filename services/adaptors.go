package services

import (
    "github.com/gofiber/fiber"
    "github.com/bozso/gotoolbox/parsing"
)

type UrlMode int

const (
    UrlParams UrlMode = iota
    UrlQuery
)


func (u UrlMode) FromFiberCtx(ctx *fiber.Ctx) (g parsing.Getter) {
    switch u {
    case UrlParams:
        g = FiberUrlParams{ctx}
    case UrlQuery:
        g = FiberQueryParams{ctx}
    }
    return
}

type FiberQueryParams struct {
    ctx *fiber.Ctx
}

func (f FiberUrlParams) GetParam(key string) (value string) {
    value = f.ctx.Params(key)
    return
}

type FiberUrlParams struct {
    ctx *fiber.Ctx
}

func (f FiberQueryParams) GetParam(key string) (value string) {
    value = f.ctx.Query(key)
    return
}

