package apiv1

import (
	"sea/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var Water = waterApi{}

type waterApi struct{}

type VerifyIDReq struct {
	InviteCode string `json:"code"`
	SeaID      string `json:"error"`
	PublicKey  string `json:"public_key"`
}

type VerifyIDResp struct {
	Random    string `json:"random"`
	PublicKey string `json:"public_key"`
}

func (*waterApi) VerifyID(r *ghttp.Request) {
	req := new(VerifyIDReq)
	r.Parse(req)

	w, _ := service.Water.GetSelfWater()

	r.Response.WriteJson(g.MapStrStr{
		"WaterId": w.WaterId,
		"SeaId":   "",
	})
}

func (*waterApi) Handler(r *ghttp.Request) {

	service.Water.ReGenWaterID()
	w, _ := service.Water.GetSelfWater()
	r.Response.WriteJson(g.MapStrStr{
		"WaterId": w.WaterId,
		"SeaId":   "",
	})
}
