package api

import (
	"sea/app/service"
	serviceWater "sea/app/service/water"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
	INVITE_RETURN_CODE_SESSION_ERROR      = 3
	INVITE_RETURN_CODE_BAD_KEY            = 4
	INVITE_RETURN_CODE_BAD_RANDOM_STRING  = 5
	INVITE_RETURN_CODE_KEY_ALREADY_EXISTS = 6
	INVITE_RETURN_CODE_SERVER_ERROR       = 7
)

func WaterInviteApiMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

}

func (*waterInviteApi) Step1(r *ghttp.Request) {
	var req *WaterInviteStep1Req
	r.Parse(req)
	throw := func(code int) {
		r.Response.WriteJson(WaterInviteStep1Resp{
			ReturnCode: code,
		})
	}
	// verify the public key
	k, err := crypto.NewKeyFromArmored(req.SenderPublicKey)
	if (err != nil) || (!k.CanVerify()) || (!k.IsPrivate()) || (!k.IsExpired()) {
		throw(INVITE_RETURN_CODE_BAD_KEY)
		return
	}
	ks := serviceWater.WaterKey.GetKeyStatus(req.SenderPublicKey)
	if (ks != serviceWater.WATER_KEY_STATUS_NOT_FOUND) || (ks == serviceWater.WATER_KEY_STATUS_WAIT_FOR_RESULT) {
		throw(INVITE_RETURN_CODE_BAD_KEY)
		return
	}
	// generate a session
	session, err := serviceWater.WaterInvite.CreateSession()
	if err != nil {
		throw(INVITE_RETURN_CODE_SESSION_ERROR)
		return
	}
	// save the public key and the session to database
	err = serviceWater.WaterInvite.SetSessionSender(session, req.SenderPublicKey)
	if err != nil {
		throw(INVITE_RETURN_CODE_SESSION_NOT_FOUND)
		return
	}
	// encrypt receiver's public key and session
	selfKeyID, err := serviceWater.WaterKey.GetSelfKeyID()
	if err != nil {
		throw(INVITE_RETURN_CODE_SERVER_ERROR)
		return
	}
	selfKey, _ := serviceWater.WaterKey.GetKey(selfKeyID)
	_ = selfKey // todo

	// fill the response
	r.Response.WriteJson(WaterInviteStep1Resp{
		Session:                    session,
		EncryptedReceiverPublicKey: "",
		ReturnCode:                 INVITE_RETURN_CODE_SUCCESS,
	})
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
