package apiv1

import (
	bgWater "sea/internal/background/water"
	"sea/internal/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var AdminWater = adminWater{}

type adminWater struct{}

type AddWaterReq struct {
	URL string `json:"url" v:"required"`
}
type AddWaterResp struct {
	URL string `json:"url"`
}

func (api *adminWater) AddWater(r *ghttp.Request) {
	req := new(AddWaterReq)
	if err := r.Parse(&req); err != nil {
		utils.ParseError(r, err)
	}
	bgWater.WaterManager.AddWater(r.Context(), "")
	g.Log().Debugf(r.Context(), "url: %s", req.URL)
	r.Response.WriteJsonExit(AddWaterResp{URL: req.URL})
}

func (api *adminWater) EditWater(r *ghttp.Request) {

}

func (api *adminWater) DeleteWater(r *ghttp.Request) {

}

func (api *adminWater) QueryWater(r *ghttp.Request) {

}
