package apiv1

import (
	"context"

	"github.com/gogf/gf/v2/os/gbuild"
)

var GetVersion = getVersionApi{}

type getVersionApi struct{}

func (*getVersionApi) BuildInfo(ctx context.Context, req *VersionReq) (*VersionRes, error) {
	return &VersionRes{
		BuildInfo: gbuild.Info(),
	}, nil
}
