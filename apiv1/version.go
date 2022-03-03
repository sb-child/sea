package apiv1

import (
	"context"

	"github.com/gogf/gf/v2/os/gbuild"
)

var GetVersion = getVersionApi{}

type getVersionApi struct{}

func (*getVersionApi) BuildInfo(ctx context.Context, req *VersionReq) (*VersionRes, error) {
	buildInfo := map[string]string{
		"commit": gbuild.Info().Git,
		"time":   gbuild.Info().Time,
		"gf":     gbuild.Info().GoFrame,
		"go":     gbuild.Info().Golang,
	}
	return &VersionRes{
		BuildInfo: buildInfo,
	}, nil
}
