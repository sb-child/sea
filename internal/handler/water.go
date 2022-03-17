package handler

import (
	"context"

	"sea/apiv1"
	"sea/internal/service"
)

var (
	Water = hWater{}
)

type hWater struct{}

func (h *hWater) Step1(ctx context.Context, req *apiv1.WaterApiJoinStep1Req) (res *apiv1.WaterApiJoinStep1Res, err error) {
	k, c, err := service.WaterJoin.JoinStep1(ctx, req.SenderPublicKey)
	return &apiv1.WaterApiJoinStep1Res{
		EncryptedReceiverPublicKey: k,
		ReturnCode:                 c,
	}, err
}

func (h *hWater) Step2(ctx context.Context, req *apiv1.WaterApiJoinStep2Req) (res *apiv1.WaterApiJoinStep2Res, err error) {
	c, err := service.WaterJoin.JoinStep2(ctx, req.EncryptedRandomString)
	return &apiv1.WaterApiJoinStep2Res{
		ReturnCode: c,
	}, err
}
