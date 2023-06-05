package fastHttpRouter

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type Handler struct {
}

func NewFastHttpHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() fasthttp.RequestHandler {
	router := routing.New()
	router.To("GET", "/", h.Hello)
	return router.HandleRequest
}
