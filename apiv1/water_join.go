package apiv1

import (
	"sea/internal/service"
	"sea/internal/utils"

	"github.com/gogf/gf/v2/net/ghttp"
)

// join api

var WaterJoin = waterJoinApi{}

type waterJoinApi struct{}

// invite steps

func (api *waterJoinApi) Step1(r *ghttp.Request) {
	req := new(WaterApiJoinStep1Req)
	if err := r.Parse(&req); err != nil {
		utils.ParseError(r, err)
	}
	k, c := service.WaterJoin.InviteStep1(r.Context(), req.SenderPublicKey)
	r.Response.WriteJsonExit(WaterApiJoinStep1Res{
		EncryptedReceiverPublicKey: k,
		ReturnCode:                 c,
	})
}
func (*waterJoinApi) Step2(r *ghttp.Request) {
	req := new(WaterApiJoinStep2Req)
	if err := r.Parse(&req); err != nil {
		utils.ParseError(r, err)
	}
	c := service.WaterJoin.InviteStep2(r.Context(), req.EncryptedRandomString)
	r.Response.WriteJson(WaterApiJoinStep2Res{
		ReturnCode: c,
	})
}
