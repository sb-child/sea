package apiv1

import (
	"context"
)

var WaterStream = waterStreamApi{}

type waterStreamApi struct{}

func (api *waterStreamApi) Transfer(ctx context.Context, req *WaterApiTransferReq) (*WaterApiTransferRes, error) {
	return &WaterApiTransferRes{
		ReturnCode: 0,
	}, nil
}
func (*waterStreamApi) Control(ctx context.Context, req *WaterApiControlReq) (*WaterApiControlRes, error) {
	return &WaterApiControlRes{
		ReturnCode: 0,
	}, nil
}
