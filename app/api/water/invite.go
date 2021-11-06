package api

import (
	// "context"
	"sea/app/service"
	// serviceWater "sea/app/service/water"

	// "github.com/ProtonMail/gopenpgp/v2/crypto"
	// "github.com/ProtonMail/gopenpgp/v2/helper"
	// "github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gparser"
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
	EncryptedReceiverPublicKey string `json:"receiver"` // a encrypted pack, sender can't be decrypted if haven't a private key
	ReturnCode                 int    `json:"returnCode"`
}

type WaterInviteStep1Pack struct {
	Session           string `json:"session"` // a 64 character random string
	ReceiverPublicKey string `json:"receiver"`
}

type WaterInviteStep2Req struct {
	EncryptedRandomString string `json:"random"` // a encrypted pack for receiver
	Session               string `json:"session"`
}

type WaterInviteStep2Resp struct {
	ReturnCode int `json:"returnCode"`
}

type WaterInviteStep2Pack struct {
	Session      string `json:"session"`
	RandomString string `json:"random"` // a 32 character random string
}

func WaterInviteApiMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

}

func (*waterInviteApi) MakeStep1Pack(session, key string) string {
	r := gparser.New(WaterInviteStep1Pack{Session: session, ReceiverPublicKey: key})
	return r.MustToJsonString()
}

// func (api *waterInviteApi) Step1(r *ghttp.Request) {

// 	api.step1(r)
// }

// func (api *waterInviteApi) step1(ctx context.Context, tx *gdb.TX, r *ghttp.Request) {
// 	var req *WaterInviteStep1Req
// 	r.Parse(req)
// 	// start a transaction

// 	// throw func
// 	throw := func(code int) {
// 		r.Response.WriteJsonExit(WaterInviteStep1Resp{
// 			ReturnCode: code,
// 		})
// 	}
// 	// verify the public key
// 	k, err := crypto.NewKeyFromArmored(req.SenderPublicKey)
// 	ks, _ := k.Armor()
// 	if (err != nil) || (serviceWater.WaterKey.CheckKey(ks, false) != serviceWater.WATER_KEY_CHECK_OK) {
// 		throw(INVITE_RETURN_CODE_BAD_KEY)
// 	}
// 	if ks := serviceWater.WaterKey.GetKeyStatus(ks); ks != serviceWater.WATER_KEY_STATUS_NOT_FOUND {
// 		throw(INVITE_RETURN_CODE_BAD_KEY)
// 	}
// 	// create a session
// 	session, err := serviceWater.WaterInvite.CreateSession()
// 	if err != nil {
// 		throw(INVITE_RETURN_CODE_SESSION_ERROR)
// 	}
// 	// save the public key and the session to database
// 	err = serviceWater.WaterInvite.SetSessionSender(session, ks)
// 	if err != nil {
// 		throw(INVITE_RETURN_CODE_SESSION_NOT_FOUND)
// 	}
// 	// encrypt receiver's public key and session
// 	selfKeyID, err := serviceWater.WaterKey.GetSelfKeyID()
// 	if err != nil {
// 		throw(INVITE_RETURN_CODE_SERVER_ERROR)
// 	}
// 	selfKey, _ := serviceWater.WaterKey.GetKey(selfKeyID) // this key has already been checked, so it's safe to continue
// 	es, err := helper.EncryptMessageArmored(selfKey, api.MakeStep1Pack(session, selfKey))
// 	if err != nil {
// 		throw(INVITE_RETURN_CODE_SERVER_ERROR)
// 	}
// 	// fill the response
// 	r.Response.WriteJsonExit(WaterInviteStep1Resp{
// 		EncryptedReceiverPublicKey: es,
// 		ReturnCode:                 INVITE_RETURN_CODE_SUCCESS,
// 	})
// }

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
