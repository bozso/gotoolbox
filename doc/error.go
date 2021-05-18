package doc

import (
	"fmt"
	"log"

	"github.com/CloudyKit/jet"
	"github.com/valyala/fasthttp"
)

type ErrorHandler interface {
	Handle(ctx *fasthttp.RequestCtx)
}

type ErrorTemplate struct {
	template *jet.Template
}

func NewErrorTemplate(template *jet.Template) (e ErrorTemplate) {
	e.template = template
	return
}

func (et ErrorTemplate) NewHandler(h Handler) (eh ErrorTemplateHandler) {
	eh.template, eh.handler = et.template, h
	return
}

type ErrorTemplateHandler struct {
	handler  Handler
	template *jet.Template
}

func (e ErrorTemplateHandler) Handle(ctx *fasthttp.RequestCtx) {
	const format = "Error while executing error template: %s\n"

	err := e.handler.Handle(ctx)
	if err == nil {
		return
	}

	log.Printf("Error: %s\n", err)

	ctx.SetContentType("text/html")
	err = e.template.Execute(ctx.Response.BodyWriter(), nil, err)

	if err != nil {
		ctx.Response.AppendBodyString(fmt.Sprintf(format, err))
	}
}
