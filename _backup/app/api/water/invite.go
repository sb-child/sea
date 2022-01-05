package api

import (
	"sea/app/service"
	serviceWater "sea/app/service/water"
	"sea/app/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// invite api
// [a water joined] -> review by admin

var WaterInvite = waterInviteApi{}

type waterInviteApi struct{}

type WaterInviteStep1Req struct {
	SenderPublicKey string `json:"sender" v:"required"` // a 4096 bits rsa public key from sender(client)
}

type WaterInviteStep1Resp struct {
	EncryptedReceiverPublicKey string `json:"receiver"` // a encrypted pack, sender can't be decrypted if haven't a private key
	ReturnCode                 int    `json:"returnCode"`
}

type WaterInviteStep2Req struct {
	EncryptedRandomString string `json:"random" v:"required"` // a encrypted pack for receiver
}

type WaterInviteStep2Resp struct {
	ReturnCode int `json:"returnCode"`
}

func WaterInviteApiMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
}

// invite steps

func (api *waterInviteApi) Step1(r *ghttp.Request) {
	req := new(WaterInviteStep1Req)
	if err := r.Parse(&req); err != nil {
		utils.ParseError(r, err)
	}
	k, c := serviceWater.WaterInvite.InviteStep1(r.Context(), req.SenderPublicKey)
	r.Response.WriteJsonExit(WaterInviteStep1Resp{
		EncryptedReceiverPublicKey: k,
		ReturnCode:                 c,
	})
}
func (*waterInviteApi) Step2(r *ghttp.Request) {
	req := new(WaterInviteStep2Req)
	if err := r.Parse(&req); err != nil {
		utils.ParseError(r, err)
	}
	c := serviceWater.WaterInvite.InviteStep2(r.Context(), req.EncryptedRandomString)
	r.Response.WriteJson(WaterInviteStep2Resp{
		ReturnCode: c,
	})
}

func (*waterInviteApi) VerifyID(r *ghttp.Request) {
	req := new(WaterInviteStep1Req)
	if err := r.Parse(&req); err != nil {
		utils.ParseError(r, err)
	}
	w, _ := service.Water.GetSelfWater()
	r.Response.WriteJson(g.MapStrStr{
		"WaterId": w.WaterId,
		"SeaId":   "",
	})
}
