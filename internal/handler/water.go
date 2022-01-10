package handler

import (
	"context"

	"sea/apiv1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Water = hWater{}
)

type hWater struct{}

func (h *hWater) Step1(ctx context.Context, req *apiv1.WaterApiJoinStep1Req) (res *apiv1.WaterApiJoinStep1Res, err error) {
	apiv1.WaterJoin.Step1(g.RequestFromCtx(ctx))
	return
}

func (h *hWater) Step2(ctx context.Context, req *apiv1.WaterApiJoinStep2Req) (res *apiv1.WaterApiJoinStep2Res, err error) {
	apiv1.WaterJoin.Step2(g.RequestFromCtx(ctx))
	return
}
