package handler

import (
	"context"

	"sea/apiv1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Admin = hAdmin{}
)

type hAdmin struct{}

func (h *hAdmin) WaterAdd(ctx context.Context, req *apiv1.AdminWaterAddReq) (res *apiv1.AdminWaterAddRes, err error) {
	apiv1.AdminWater.AddWater(g.RequestFromCtx(ctx))
	return
}

func (h *hAdmin) WaterDelete(ctx context.Context, req *apiv1.AdminWaterDeleteReq) (res *apiv1.AdminWaterDeleteRes, err error) {
	apiv1.AdminWater.DeleteWater(g.RequestFromCtx(ctx))
	return
}

func (h *hAdmin) WaterEdit(ctx context.Context, req *apiv1.AdminWaterEditReq) (res *apiv1.AdminWaterEditRes, err error) {
	apiv1.AdminWater.EditWater(g.RequestFromCtx(ctx))
	return
}

func (h *hAdmin) WaterQuery(ctx context.Context, req *apiv1.AdminWaterQueryReq) (res *apiv1.AdminWaterQueryRes, err error) {
	apiv1.AdminWater.QueryWater(g.RequestFromCtx(ctx))
	return
}
