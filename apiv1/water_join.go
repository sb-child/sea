package apiv1

import (
	"context"
	"sea/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// join api

var WaterJoin = waterJoinApi{}

type waterJoinApi struct{}

// invite steps

func (api *waterJoinApi) Step1(ctx context.Context, req *WaterApiJoinStep1Req) (*WaterApiJoinStep1Res, error) {
	g.Log().Debugf(ctx, "%v", req)
	k, c, err := service.WaterJoin.InviteStep1(ctx, req.SenderPublicKey)
	return &WaterApiJoinStep1Res{
		EncryptedReceiverPublicKey: k,
		ReturnCode:                 c,
	}, err
}
func (*waterJoinApi) Step2(ctx context.Context, req *WaterApiJoinStep2Req) (*WaterApiJoinStep2Res, error) {
	c, err := service.WaterJoin.InviteStep2(ctx, req.EncryptedRandomString)
	return &WaterApiJoinStep2Res{
		ReturnCode: c,
	}, err
}
