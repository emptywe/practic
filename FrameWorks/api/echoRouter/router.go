package echoRouter

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
}

func NewEchoHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() http.Handler {
	router := echo.New()
	router.Add("GET", "/", h.Hello)
	return router
}
