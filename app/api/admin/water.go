package api

import (
	bgWater "sea/app/background/water"

	"github.com/gogf/gf/v2/net/ghttp"
)

var AdminWater = adminWater{}

type adminWater struct{}

func (api *adminWater) AddWater(r *ghttp.Request) {
	bgWater.WaterManager.AddWater(r.Context(), "")
}

func (api *adminWater) EditWater(r *ghttp.Request) {

}

func (api *adminWater) DeleteWater(r *ghttp.Request) {

}

func (api *adminWater) QueryWater(r *ghttp.Request) {

}
