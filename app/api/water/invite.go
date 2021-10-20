package api

import (
	"sea/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// invite api
// [a water joined] -> review by admin

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
	EncryptedRandomString string `json:"random"` // a 32 character random string
	Session               string `json:"session"`
}

type WaterInviteStep2Resp struct {
	ReturnCode int `json:"returnCode"`
}

const (
	INVITE_RETURN_CODE_SUCCESS           = 0
	INVITE_RETURN_CODE_DECRYPTION_FAILED = 1
	INVITE_RETURN_CODE_SESSION_NOT_FOUND = 2
	INVITE_RETURN_CODE_BAD_KEY           = 3
	INVITE_RETURN_CODE_BAD_RANDOM_STRING = 4
	INVITE_RETURN_CODE_ALREADY_EXISTS    = 5
)

func WaterInviteApiMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

}

func (*waterInviteApi) Step1(r *ghttp.Request) {
	var req *WaterInviteStep1Req
	r.Parse(req)

	// w, _ := service.Water.GetSelfWater()

	r.Response.WriteJson(WaterInviteStep1Resp{})
}

func (*waterInviteApi) Step2(r *ghttp.Request) {
	var req *WaterInviteStep2Req
	r.Parse(req)

	// w, _ := service.Water.GetSelfWater()

	r.Response.WriteJson(WaterInviteStep2Resp{})
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
