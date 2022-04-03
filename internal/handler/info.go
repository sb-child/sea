package handler

import (
	"context"

	"sea/apiv1"

	"sea/internal/service"

	"github.com/gogf/gf/v2/os/gbuild"
)

var (
	Version = hVersion{}
	Info    = hInfo{}
)

type hVersion struct{}
type hInfo struct{}

func (h *hVersion) GetVersion(ctx context.Context, req *apiv1.VersionReq) (res *apiv1.VersionRes, err error) {
	buildInfo := map[string]string{
		"commit": gbuild.Info().Git,
		"time":   gbuild.Info().Time,
		"gf":     gbuild.Info().GoFrame,
		"go":     gbuild.Info().Golang,
	}
	return &apiv1.VersionRes{
		BuildInfo: buildInfo,
	}, nil
}

func (h *hVersion) GetPublicKey(ctx context.Context, req *apiv1.PublicKeyReq) (res *apiv1.PublicKeyRes, err error) {
	w, err := service.Water.GetSelfWater()
	if err != nil {
		return nil, err
	}
	k, _ := service.UnpackPrivateKey(w.WaterKey)
	k.N.Bytes()
	return &apiv1.PublicKeyRes{
		Key: w.WaterKey,
		Id:  w.WaterId,
	}, nil
}
