package fastHttpRouter

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

func (h *Handler) Hello(ctx *routing.Context) error {
	_, err := ctx.Write([]byte("hello i'm fasthttp router"))
	return err
}
