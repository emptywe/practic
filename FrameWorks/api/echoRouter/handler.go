package echoRouter

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello i'm echo router")
}
