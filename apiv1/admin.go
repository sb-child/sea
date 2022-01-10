package apiv1

import (
	"context"
	bgWater "sea/internal/background/water"

	"github.com/gogf/gf/v2/frame/g"
)

var AdminWater = adminWater{}

type adminWater struct{}

func (api *adminWater) AddWater(ctx context.Context, req *AdminWaterAddReq) (*AdminWaterAddRes, error) {
	bgWater.WaterManager.AddWater(ctx, req.URL)
	g.Log().Debugf(ctx, "url: %s", req.URL)
	return &AdminWaterAddRes{ReturnCode: 0}, nil
}

func (api *adminWater) EditWater(ctx context.Context, req *AdminWaterEditReq) (*AdminWaterEditRes, error) {
	return &AdminWaterEditRes{}, nil
}

func (api *adminWater) DeleteWater(ctx context.Context, req *AdminWaterDeleteReq) (*AdminWaterDeleteRes, error) {
	return &AdminWaterDeleteRes{}, nil
}

func (api *adminWater) QueryWater(ctx context.Context, req *AdminWaterQueryReq) (*AdminWaterQueryRes, error) {
	return &AdminWaterQueryRes{}, nil
}
