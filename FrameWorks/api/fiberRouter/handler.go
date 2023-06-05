package fiberRouter

import "github.com/gofiber/fiber/v2"

func (h *Handler) Hello(ctx *fiber.Ctx) error {
	_, err := ctx.Write([]byte("hello i'm fiber router"))
	return err
}
