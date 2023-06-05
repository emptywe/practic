package ginRouter

import "github.com/gin-gonic/gin"

func (h *Handler) Hello(ctx *gin.Context) {
	ctx.Writer.Write([]byte("hello i'm gin router"))
}
