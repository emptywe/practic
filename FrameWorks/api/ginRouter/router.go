package ginRouter

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewGinHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Handle("GET", "/", h.Hello)
	return router
}
