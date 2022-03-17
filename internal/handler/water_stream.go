package handler

import (
	"context"

	"sea/apiv1"
)

var (
	WaterStream = hWaterStream{}
)

type hWaterStream struct{}

func (h *hWaterStream) Transfer(ctx context.Context, req *apiv1.WaterApiTransferReq) (res *apiv1.WaterApiTransferRes, err error) {
	res, err = &apiv1.WaterApiTransferRes{}, nil
	return
}

func (h *hWaterStream) Control(ctx context.Context, req *apiv1.WaterApiControlReq) (res *apiv1.WaterApiControlRes, err error) {
	res, err = &apiv1.WaterApiControlRes{}, nil
	return
}
