package apiv1

import (
	"context"
	bgWater "sea/internal/background/water"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var AdminWater = adminWater{}

type adminWater struct{}

func (api *adminWater) AddWater(ctx context.Context, req *AdminWaterAddReq) (*AdminWaterAddRes, error) {
	bgWater.WaterManager.AddWater(ctx, req.URL)
	g.Log().Debugf(ctx, "url: %s", req.URL)
	return &AdminWaterAddRes{ReturnCode: 0}, nil
}

func (api *adminWater) EditWater(r *ghttp.Request) {

}

func (api *adminWater) DeleteWater(r *ghttp.Request) {

}

func (api *adminWater) QueryWater(r *ghttp.Request) {

}
