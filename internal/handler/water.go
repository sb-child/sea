package handler

import (
	"context"

	"sea/apiv1"
	waterApi "sea/apiv1/water"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Water = hWater{}
)

type hWater struct{}

func (h *hWater) Step1(ctx context.Context, req *apiv1.WaterApiInviteStep1Req) (res *apiv1.WaterApiInviteStep1Res, err error) {
	waterApi.WaterInvite.Step1(g.RequestFromCtx(ctx))
	return
}

func (h *hWater) Step2(ctx context.Context, req *apiv1.WaterApiInviteStep2Req) (res *apiv1.WaterApiInviteStep2Res, err error) {
	waterApi.WaterInvite.Step2(g.RequestFromCtx(ctx))
	return
}
