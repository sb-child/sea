package api

import (
	"sea/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var WaterInvite = waterInviteApi{}

type waterInviteApi struct{}

type WaterInviteStep1Req struct {
	SenderPublicKey string `json:"sender"`
	Session         string `json:"session"`
}

type WaterInviteStep1Resp struct {
	ReturnCode                 int    `json:"returnCode"`
	EncryptedReceiverPublicKey string `json:"receiver"`
}

type WaterInviteStep2Req struct {
	EncryptedRandomString string `json:"random"`
	Session               string `json:"session"`
}

type WaterInviteStep2Resp struct {
	ReturnCode int `json:"returnCode"`
}

const (
	WATER_INVITE_RETURN_CODE_SUCCESS = 0
	WATER_INVITE_RETURN_CODE_DECRYPTION_FAILED = 1
	WATER_INVITE_RETURN_CODE_SESSION_NOT_FOUND = 2
)

func WaerInviteApiMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

}

func (*waterInviteApi) VerifyID(r *ghttp.Request) {
	var req *WaterInviteStep1Req
	r.Parse(req)

	w, _ := service.Water.GetSelfWater()

	r.Response.WriteJson(g.MapStrStr{
		"WaterId": w.WaterId,
		"SeaId":   "",
	})
}

func (*waterInviteApi) Handler(r *ghttp.Request) {

	service.Water.ReGenWaterID()
	w, _ := service.Water.GetSelfWater()
	r.Response.WriteJson(g.MapStrStr{
		"WaterId": w.WaterId,
		"SeaId":   "",
	})
}
