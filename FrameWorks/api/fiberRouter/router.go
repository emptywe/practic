package fiberRouter

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func NewFiberHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutesAndServer(host, port string) error {

	router := fiber.New()

	router.Add("GET", "/", h.Hello)

	return router.Listen(host + ":" + port)
}
