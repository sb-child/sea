package api

import (
	bgWater "sea/app/background/water"
	"sea/app/utils"

	"github.com/gogf/gf/v2/net/ghttp"
)

var AdminWater = adminWater{}

type adminWater struct{}

type AddWaterReq struct {
	URL string `json:"url" v:"required"`
}

func (api *adminWater) AddWater(r *ghttp.Request) {
	req := new(AddWaterReq)
	if err := r.Parse(&req); err != nil {
		utils.ParseError(r, err)
	}
	bgWater.WaterManager.AddWater(r.Context(), "")
}

func (api *adminWater) EditWater(r *ghttp.Request) {

}

func (api *adminWater) DeleteWater(r *ghttp.Request) {

}

func (api *adminWater) QueryWater(r *ghttp.Request) {

}
