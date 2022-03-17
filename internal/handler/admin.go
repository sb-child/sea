package handler

import (
	"context"
	"sea/apiv1"
	bgWater "sea/internal/background/water"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Admin = hAdmin{}
)

type hAdmin struct{}

func (h *hAdmin) WaterAdd(ctx context.Context, req *apiv1.AdminWaterAddReq) (res *apiv1.AdminWaterAddRes, err error) {
	bgWater.WaterManager.AddWater(ctx, req.URL)
	g.Log().Debugf(ctx, "url: %s", req.URL)
	res, err = &apiv1.AdminWaterAddRes{ReturnCode: 0}, nil
	return
}

func (h *hAdmin) WaterDelete(ctx context.Context, req *apiv1.AdminWaterDeleteReq) (res *apiv1.AdminWaterDeleteRes, err error) {
	res, err = &apiv1.AdminWaterDeleteRes{}, nil
	return
}

func (h *hAdmin) WaterEdit(ctx context.Context, req *apiv1.AdminWaterEditReq) (res *apiv1.AdminWaterEditRes, err error) {
	res, err = &apiv1.AdminWaterEditRes{}, nil
	return
}

func (h *hAdmin) WaterQuery(ctx context.Context, req *apiv1.AdminWaterQueryReq) (res *apiv1.AdminWaterQueryRes, err error) {
	res, err = &apiv1.AdminWaterQueryRes{}, nil
	return
}
