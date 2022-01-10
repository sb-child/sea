package handler

import (
	"context"

	"sea/apiv1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Version = hVersion{}
)

type hVersion struct{}

func (h *hVersion) GetVersion(ctx context.Context, req *apiv1.VersionReq) (res *apiv1.VersionRes, err error) {
	res, err = apiv1.GetVersion.BuildInfo(ctx, req)
	g.Log().Debug(ctx, "version", res)
	return
}
