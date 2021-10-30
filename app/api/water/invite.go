package api

import (
	"sea/app/service"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// invite api
// [a water joined] -> review by admin

var WaterInvite = waterInviteApi{}

type waterInviteApi struct{}

type WaterInviteStep1Req struct {
	SenderPublicKey string `json:"sender"` // a 4096 bits OpenPGP public key
}

type WaterInviteStep1Resp struct {
	Session                    string `json:"session"`  // a 64 character random string
	EncryptedReceiverPublicKey string `json:"receiver"` // a encrypted OpenPGP public key
	ReturnCode                 int    `json:"returnCode"`
}

type WaterInviteStep2Req struct {
	EncryptedRandomString string `json:"random"` // a 32 character random string
	Session               string `json:"session"`
}

type WaterInviteStep2Resp struct {
	ReturnCode int `json:"returnCode"`
}

const (
	INVITE_RETURN_CODE_SUCCESS            = 0
	INVITE_RETURN_CODE_DECRYPTION_FAILED  = 1
	INVITE_RETURN_CODE_SESSION_NOT_FOUND  = 2
	INVITE_RETURN_CODE_BAD_KEY            = 3
	INVITE_RETURN_CODE_BAD_RANDOM_STRING  = 4
	INVITE_RETURN_CODE_KEY_ALREADY_EXISTS = 5
)

func WaterInviteApiMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

}

func (*waterInviteApi) Step1(r *ghttp.Request) {
	var req *WaterInviteStep1Req
	r.Parse(req)
	// verify the public key
	k, err := crypto.NewKeyFromArmored(req.SenderPublicKey)
	if (err != nil) || (!k.CanVerify()) || (!k.IsPrivate()) || (!k.IsExpired()) {
		r.Response.WriteJson(WaterInviteStep1Resp{
			ReturnCode: INVITE_RETURN_CODE_BAD_KEY,
		})
		return
	}
	// save the public key to database

	// generate a session

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
