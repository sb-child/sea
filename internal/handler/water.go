package handler

import (
	"context"

	"sea/apiv1"
)

var (
	Water = hWater{}
)

type hWater struct{}

func (h *hWater) Step1(ctx context.Context, req *apiv1.WaterApiJoinStep1Req) (res *apiv1.WaterApiJoinStep1Res, err error) {
	res, err = apiv1.WaterJoin.Step1(ctx, req)
	return
}

func (h *hWater) Step2(ctx context.Context, req *apiv1.WaterApiJoinStep2Req) (res *apiv1.WaterApiJoinStep2Res, err error) {
	res, err = apiv1.WaterJoin.Step2(ctx, req)
	return
}

func (h *hWater) Transfer(ctx context.Context, req *apiv1.WaterApiTransferReq) (res *apiv1.WaterApiTransferRes, err error) {
	res, err = apiv1.WaterStream.Transfer(ctx, req)
	return
}

func (h *hWater) Control(ctx context.Context, req *apiv1.WaterApiControlReq) (res *apiv1.WaterApiControlRes, err error) {
	res, err = apiv1.WaterStream.Control(ctx, req)
	return
}
