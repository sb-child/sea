package handler

import (
	"context"

	"sea/apiv1"
)

var (
	Admin = hAdmin{}
)

type hAdmin struct{}

func (h *hAdmin) WaterAdd(ctx context.Context, req *apiv1.AdminWaterAddReq) (res *apiv1.AdminWaterAddRes, err error) {
	res, err = apiv1.AdminWater.AddWater(ctx, req)
	return
}

func (h *hAdmin) WaterDelete(ctx context.Context, req *apiv1.AdminWaterDeleteReq) (res *apiv1.AdminWaterDeleteRes, err error) {
	res, err = apiv1.AdminWater.DeleteWater(ctx, req)
	return
}

func (h *hAdmin) WaterEdit(ctx context.Context, req *apiv1.AdminWaterEditReq) (res *apiv1.AdminWaterEditRes, err error) {
	res, err = apiv1.AdminWater.EditWater(ctx, req)
	return
}

func (h *hAdmin) WaterQuery(ctx context.Context, req *apiv1.AdminWaterQueryReq) (res *apiv1.AdminWaterQueryRes, err error) {
	res, err = apiv1.AdminWater.QueryWater(ctx, req)
	return
}
