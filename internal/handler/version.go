package handler

import (
	"context"

	"sea/apiv1"

	"github.com/gogf/gf/v2/os/gbuild"
)

var (
	Version = hVersion{}
)

type hVersion struct{}

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
