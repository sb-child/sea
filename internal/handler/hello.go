package handler

import (
	"context"

	"sea/apiv1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Hello = hHello{}
)

type hHello struct{}

func (h *hHello) Hello(ctx context.Context, req *apiv1.HelloReq) (res *apiv1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
